package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Consumer interface {
	Consume(ctx context.Context, l *zap.SugaredLogger, topic string, ch chan<- *sarama.ConsumerMessage) error
}

type SaramaConsumer struct {
	group sarama.ConsumerGroup
}

func NewConsumer(config *Config, consumerGroup string) (*SaramaConsumer, error) {
	c := sarama.NewConfig()

	c.Consumer.Return.Errors = true
	c.Net.SASL.Enable = config.UseAuthentication
	c.Net.SASL.User = config.Username
	c.Net.SASL.Password = config.Password

	group, err := sarama.NewConsumerGroup(config.Addresses, consumerGroup, c)
	if err != nil {
		return nil, err
	}

	return &SaramaConsumer{
		group: group,
	}, nil
}

func (c *SaramaConsumer) Consume(ctx context.Context, l *zap.SugaredLogger, topic string, msgCh chan<- *sarama.ConsumerMessage, ackCh chan bool) error {
	messageHandler := newConsumerGroupHandler(msgCh, ackCh)
	defer close(msgCh)
	for {
		select {
		case <-ctx.Done():
			l.Infow("receive stop signal", "topic", topic)
			return nil
		default:
			err := c.group.Consume(ctx, []string{topic}, messageHandler)
			if err != nil {
				l.Errorw("error when consume messages", "topic", topic, "error", err)
				return err
			}
		}
	}
}

type consumerGroupHandler struct {
	msgChan chan<- *sarama.ConsumerMessage
	ackChan chan bool
}

func newConsumerGroupHandler(msgChan chan<- *sarama.ConsumerMessage, ackCh chan bool) *consumerGroupHandler {
	return &consumerGroupHandler{
		msgChan: msgChan,
		ackChan: ackCh,
	}
}

func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim consumes messages from the claim and pushes them to the channel.
// It manually commits the offset once the message is processed.
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// Push the message to the channel
		h.msgChan <- msg

		ack := <-h.ackChan
		if ack {
			// Acknowledge the message by marking it as processed (commit the offset)
			session.MarkMessage(msg, "")
		} else {
			break // stop consume when fail to broadcast message
		}
	}
	return nil
}
