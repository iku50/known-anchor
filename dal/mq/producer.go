package mq

import (
	"context"
	"fmt"
	"known-anchors/config"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer  *kafka.Writer
	Topic   string
	ProChan chan Message
}

func NewProducer(topic string) *Producer {
	return &Producer{
		Writer: &kafka.Writer{
			Addr:                   kafka.TCP(config.Conf.Kafka.Addr),
			Topic:                  topic,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
		Topic:   topic,
		ProChan: make(chan Message, 100),
	}
}

func (p *Producer) Produce() {
	for msg := range p.ProChan {
		if err := p.Writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(msg.Key),
			Value: []byte(msg.Value),
		}); err != nil {
			fmt.Println("failed to write messages:", err)
		}
	}
}

func (p *Producer) Close() error {
	if err := p.Writer.Close(); err != nil {
		return err
	}
	return nil
}
