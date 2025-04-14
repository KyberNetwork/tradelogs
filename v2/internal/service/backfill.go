package service

import (
	"fmt"

	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"go.uber.org/zap"
)

type Backfill struct {
	storage backfill.IStorage
	l       *zap.SugaredLogger
	worker  *worker.BackFiller
}

const (
	MaxBackfillTaskNumber = 10
	SystemBackfillTaskID  = -1
)

func NewBackfillService(storage backfill.IStorage, l *zap.SugaredLogger, w *worker.BackFiller) (*Backfill, error) {
	srv := &Backfill{
		storage: storage,
		l:       l,
		worker:  w,
	}
	err := srv.rerunAllTasks()
	if err != nil {
		return nil, fmt.Errorf("fail to rerun all tasks: %w", err)
	}
	return srv, nil
}

func (s *Backfill) rerunAllTasks() error {
	tasks, err := s.storage.GetTask()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.Status == backfill.StatusTypeProcessing || task.Status == backfill.StatusTypeFailed {
			go s.worker.BackfillByTask(task)
		}
	}
	return nil
}

func (s *Backfill) NewBackfillForCowTrade(from, to uint64) (int, string, error) {
	if from > to {
		return 0, "", fmt.Errorf("startBlock must be lower than endBlock")
	}

	task := backfill.Task{
		FromBlock: from,
		ToBlock:   to,
		Exchange:  constant.CowProtocol,
	}
	id, err := s.storage.CreateTask(task)
	if err != nil {
		return 0, "", fmt.Errorf("cannot create backfill task: %w", err)
	}

	task.ID = id
	go s.worker.BackfillForCowProtocol(task)

	return id, "", nil
}

func (s *Backfill) NewBackfillForTradeLogs(from, to uint64, exchange string) (int, string, error) {
	var message string

	if !s.worker.IsValidExchange(exchange) {
		return 0, "", fmt.Errorf("invalid exchange %s", exchange)
	}

	first, last, _, err := s.worker.GetBlockRanges()
	if err != nil {
		return 0, "", fmt.Errorf("cannot get ongoing backfill block ranges: %w", err)
	}

	// new task block range will be processed by the default backfill flow
	if from >= first && to <= last {
		return 0, "", fmt.Errorf("new range covered by ongoing backfill process")
	}

	// disable backfill the older block to simplify the logic
	if from < first {
		return 0, "", fmt.Errorf("from block cannot smaller than the deployed block")
	}

	// new task block range covered partially, set the new range to avoid duplication
	if from < last {
		from = last
		message = fmt.Sprintf("new range partially covered by ongoing backfill process, changed old from_block %d to %d", from, last)
	}

	task := backfill.Task{
		FromBlock: from,
		ToBlock:   to,
		Exchange:  exchange,
	}
	id, err := s.storage.CreateTask(task)
	if err != nil {
		return 0, "", fmt.Errorf("cannot create backfill task: %w", err)
	}

	task.ID = id
	go s.worker.BackfillByExchange(task)

	return id, message, nil
}

func (s *Backfill) NewBackfillTask(from, to uint64, exchange string) (int, string, error) {
	// limit max 10 tasks running at the same time
	count, err := s.storage.GetRunningTaskNumber()
	if err != nil {
		return 0, "", fmt.Errorf("fail to get running task number: %w", err)
	}
	if count >= MaxBackfillTaskNumber {
		return 0, "", fmt.Errorf("number of running task exceed: %d", MaxBackfillTaskNumber)
	}
	switch exchange {
	case constant.CowProtocol:
		id, message, err := s.NewBackfillForCowTrade(from, to)
		return id, message, err
	default:
		id, message, err := s.NewBackfillForTradeLogs(from, to, exchange)
		return id, message, err
	}
}

func (s *Backfill) CancelBackfillTask(id int) error {
	if id == SystemBackfillTaskID {
		return s.worker.CancelSystemBackfill()
	}
	task, err := s.storage.GetTaskByID(id)
	if err != nil {
		return fmt.Errorf("cannot get backfill task with id %d: %w", id, err)
	}
	if task.Status != backfill.StatusTypeProcessing {
		return fmt.Errorf("task with id %d is not processing, current status: %s", id, task.Status)
	}
	err = s.storage.UpdateTask(task.ID, nil, backfill.StatusTypeCanceled)
	if err != nil {
		return fmt.Errorf("cannot cancel backfill task with id %d: %w", id, err)
	}
	return nil
}

func (s *Backfill) RestartBackfillTask(id int) error {
	if id == SystemBackfillTaskID {
		return s.worker.RunSystemBackfill()
	}
	task, err := s.storage.GetTaskByID(id)
	if err != nil {
		return fmt.Errorf("cannot get backfill task with id %d: %w", id, err)
	}

	// limit max 10 tasks running at the same time
	count, err := s.storage.GetRunningTaskNumber()
	if err != nil {
		return fmt.Errorf("fail to get running task number: %w", err)
	}
	if count >= MaxBackfillTaskNumber {
		return fmt.Errorf("number of running task exceed: %d", MaxBackfillTaskNumber)
	}

	if task.Status == backfill.StatusTypeProcessing || task.Status == backfill.StatusTypeDone {
		return fmt.Errorf("cannot restart task with id %d, current status: %s", id, task.Status)
	}

	go s.worker.BackfillByExchange(task)
	return nil
}

func (s *Backfill) ListTask() ([]backfill.Task, error) {
	tasks, err := s.storage.GetTask()
	if err != nil {
		return nil, err
	}
	// system backfill task
	first, last, _, err := s.worker.GetBlockRanges()
	if err != nil {
		return tasks, nil
	}
	status := backfill.StatusTypeProcessing
	if first >= last {
		status = backfill.StatusTypeDone
	}
	task := backfill.Task{
		ID:             SystemBackfillTaskID,
		FromBlock:      first,
		ToBlock:        last,
		ProcessedBlock: last,
		Status:         status,
	}
	tasks = append([]backfill.Task{task}, tasks...)
	return tasks, nil
}

func (s *Backfill) GetTask(id int) (backfill.Task, error) {
	return s.storage.GetTaskByID(id)
}
