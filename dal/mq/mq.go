package mq

import "context"

type MQMessage struct {
	Topic string
	Body  []byte
}

type Producer interface {
	Produce(ctx context.Context, message *MQMessage) error
}

type Consumer interface {
	Consume(ctx context.Context, topic string) (*MQMessage, error)
}

type MQ interface {
	Producer
	Consumer
}
