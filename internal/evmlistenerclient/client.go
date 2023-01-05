package evmlistenerclient

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const (
	messageKey = "message"
)

var (
	blockTime = 12 * time.Second
)

type Config struct {
	Topic         string
	GroupName     string
	GroupConsumer string
}

type Client struct {
	client     redis.UniversalClient
	config     Config
	l          *zap.SugaredLogger
	isInitDone bool
	groupName  string
}

func New(l *zap.SugaredLogger, cfg Config, client redis.UniversalClient) *Client {
	return &Client{
		client:    client,
		config:    cfg,
		l:         l,
		groupName: "trading-tradelogs",
	}
}

func (c *Client) Init(ctx context.Context) error {
	_, err := c.client.XGroupCreate(ctx, c.config.Topic, c.groupName, "0").Result()
	if err != nil {
		if err.Error() == "BUSYGROUP Consumer Group name already exists" {
			return nil
		}
		c.l.Errorw("Init group", "err", err)
		return err
	}
	return err
}

func (c *Client) GConsume(ctx context.Context) ([]Message, error) {
	var (
		messages []Message
		id       = ">"
	)
	if !c.isInitDone {
		id = "0"
	}
	c.l.Debugw("Consuming msg", "isInitDone", c.isInitDone)
	streams, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Streams:  []string{c.config.Topic, id},
		Group:    c.config.GroupName,
		Consumer: c.config.GroupConsumer,
		Block:    time.Second,
	}).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		c.l.Errorw("XREADGROUP", "error", err)
		return nil, err
	}

	if !c.isInitDone && len(streams[0].Messages) == 0 {
		c.isInitDone = true
		return messages, nil
	}

	for _, msg := range streams[0].Messages {
		newMessage := Message{}
		newMessage.ID = msg.ID
		if data, ok := msg.Values[messageKey]; ok {
			err := json.Unmarshal([]byte(data.(string)), &newMessage)
			if err != nil {
				c.l.Errorf("error on unmarshal stream message:%v\n", msg.ID)
			}
		}

		messages = append(messages, newMessage)
	}
	return messages, nil
}

func (c *Client) Ack(ctx context.Context, m []Message) error {
	if len(m) == 0 {
		return nil
	}
	var ids []string
	for _, message := range m {
		ids = append(ids, message.ID)
	}
	_, err := c.client.XAck(ctx, c.config.Topic, c.groupName, ids...).Result()
	if err != nil {
		c.l.Errorw("XAck", "ids", ids, "error", err)
		return err
	}
	return nil
}

func (c *Client) Consume(
	ctx context.Context,
	id string,
) ([]types.Message, error) {
	streams, err := c.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{c.config.Topic, id},
		Count:   1,
		Block:   blockTime,
	}).Result()

	if err != nil {
		c.l.Errorw("err on consume events", "err", err)
		return nil, err
	}

	var messages = make([]types.Message, 0, len(streams[0].Messages))
	for _, msg := range streams[0].Messages {
		newMessage := types.Message{}

		err := json.Unmarshal([]byte(msg.Values[messageKey].(string)), &newMessage)
		if err != nil {
			c.l.Errorf("error on unmarshal stream message:%v\n", msg.ID)
			continue
		}
		messages = append(messages, newMessage)
	}

	return messages, nil
}
