package backfill

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IStorage interface {
	Get() ([]State, error)
	Update(backfilled uint64) error

	CreateTask(task Task) (int, error)
	UpdateTask(id int, block *uint64, status string) error
	GetTaskByID(taskID int) (Task, error)
	GetTask() ([]Task, error)
	DeleteTask(taskID int) error
	GetRunningTaskNumber() (int, error)
}

type Storage struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

type (
	State struct {
		Exchange        string `db:"exchange" json:"exchange"`
		DeployBlock     uint64 `db:"deploy_block" json:"deploy_block"`
		BackFilledBlock uint64 `db:"backfilled_block" json:"backfilled_block"`
	}
	Task struct {
		ID             int       `db:"id" json:"id"`
		Exchange       string    `db:"exchange" json:"exchange"`
		FromBlock      uint64    `db:"from_block" json:"from_block"`
		ToBlock        uint64    `db:"to_block" json:"to_block"`
		ProcessedBlock uint64    `db:"processed_block" json:"processed_block"`
		CreatedAt      time.Time `db:"created_at" json:"created_at,omitempty"`
		UpdatedAt      time.Time `db:"updated_at" json:"updated_at,omitempty"`
		Status         string    `db:"status" json:"status"`
	}
)

const (
	StatusTypeProcessing = "processing"
	StatusTypeFailed     = "failed"
	StatusTypeDone       = "done"
	StatusTypeCanceled   = "canceled"
)

func New(l *zap.SugaredLogger, db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		l:  l,
	}
}

func (s *Storage) Get() ([]State, error) {
	var infos []State
	err := s.db.Select(&infos, "SELECT * FROM backfill ORDER BY deploy_block ASC")
	if err != nil {
		s.l.Errorw("failed to fetch all backfill", "err", err)
		return nil, fmt.Errorf("failed to fetch all backfill state: %w", err)
	}
	return infos, nil
}

func (s *Storage) Update(backfilled uint64) error {
	res, err := s.db.Exec("UPDATE backfill SET backfilled_block=$1 WHERE backfilled_block > $2 OR backfilled_block = 0", backfilled, backfilled)
	if err != nil {
		s.l.Errorw("failed to update backfill", "err", err)
		return fmt.Errorf("failed to update backfill: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.l.Errorw("failed to get rows affected", "err", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	s.l.Infow("backfill updated", "rowsAffected", rowsAffected)
	return nil
}

func (s *Storage) CreateTask(task Task) (int, error) {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("backfill_task").
		Columns([]string{"exchange", "from_block", "to_block"}...).
		Values(task.Exchange, task.FromBlock, task.ToBlock).
		Suffix("RETURNING id")
	q, p, err := b.ToSql()
	if err != nil {
		s.l.Errorw("fail to build insert task statement", "error", err)
		return 0, fmt.Errorf("fail to build insert task statement: %w", err)
	}
	var id int
	if err = s.db.QueryRow(q, p...).Scan(&id); err != nil {
		s.l.Errorw("fail to exec insert backfill task", "sql", q, "arg", p, "error", err)
		return 0, fmt.Errorf("fail to exec insert backfill task: %w", err)
	}
	return id, nil
}

func (s *Storage) UpdateTask(id int, block *uint64, status string) error {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update("backfill_task").
		Set("updated_at", time.Now())

	if block != nil {
		b = b.Set("processed_block", *block)
	}
	if len(status) != 0 {
		b = b.Set("status", status)
	}
	b = b.Where(squirrel.Eq{"id": id})

	q, p, err := b.ToSql()
	if err != nil {
		s.l.Errorw("fail to build update task statement", "error", err)
		return fmt.Errorf("failed to build update task statement: %w", err)
	}

	_, err = s.db.Exec(q, p...)
	if err != nil {
		s.l.Errorw("failed to update backfill task", "err", err, "id", id)
		return fmt.Errorf("failed to update backfill task: %w", err)
	}
	return nil
}

func (s *Storage) GetTaskByID(taskID int) (Task, error) {
	var tasks []Task
	err := s.db.Select(&tasks, "SELECT * FROM backfill_task WHERE id=$1", taskID)
	if err != nil {
		s.l.Errorw("failed to fetch backfill task", "err", err, "id", taskID)
		return Task{}, fmt.Errorf("failed to fetch all backfill task: %w", err)
	}
	if len(tasks) != 1 {
		s.l.Errorw("invalid result length", "id", taskID, "length", len(tasks))
		return Task{}, fmt.Errorf("cannot query task with id %d", taskID)
	}
	return tasks[0], nil
}

func (s *Storage) GetTask() ([]Task, error) {
	var task []Task
	err := s.db.Select(&task, "SELECT * FROM backfill_task")
	if err != nil {
		s.l.Errorw("failed to fetch all backfill", "err", err)
		return nil, fmt.Errorf("failed to fetch all backfill task: %w", err)
	}
	return task, nil
}

func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Exec("DELETE FROM backfill_task WHERE id=$1", taskID)
	if err != nil {
		s.l.Errorw("failed to delete backfill task", "err", err, "id", taskID)
		return fmt.Errorf("failed to delete backfill task: %w", err)
	}
	return nil
}

func (s *Storage) GetRunningTaskNumber() (int, error) {
	var count int
	err := s.db.Get(&count, "SELECT count(*) FROM backfill_task WHERE status = $1", StatusTypeProcessing)
	if err != nil {
		s.l.Errorw("failed to count backfill", "err", err)
		return 0, fmt.Errorf("failed to count backfill task: %w", err)
	}
	return count, nil
}
