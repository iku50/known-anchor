package messages

import (
	"context"
	"known-anchors/dal/mq"
)

type MessageService struct {
	MQ    mq.MQ
	Queue string
}

func NewMessageService(mq mq.MQ, queue string) *MessageService {
	return &MessageService{
		MQ:    mq,
		Queue: queue,
	}
}

func (s *MessageService) ProduceMessage(ctx context.Context, message *mq.MQMessage) error {
	return s.MQ.Produce(ctx, message)
}

func (s *MessageService) ConsumeMessage(ctx context.Context) (*mq.MQMessage, error) {
	for {
		message, err := s.MQ.Consume(ctx, s.Queue)
		if err != nil {
			return nil, err
		}
		if message != nil {
			switch message.Topic {
			case "email":
				// todo: 从 topic 中解析出 email
				
			}
		}
	}
}
