package mail

import (
	"context"
	"known-anchors/config"
	"known-anchors/util/pool"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type MailConsumer struct {
	Reader *kafka.Reader
	Pool   *pool.Pool
}

func NewMailConsumer() *MailConsumer {
	return &MailConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{config.Conf.Kafka.Addr},
			Topic:    "mail",
			GroupID:  "MailConsumerGroupID",
			MaxBytes: 10e6, // 10MB
		}),
		Pool: pool.NewPool(10),
	}
}

func (c *MailConsumer) Consume() {
	for {
		m, err := c.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
			break
		}
		c.Pool.Submit(func() {
			mail, err := JsonToMailCode(string(m.Value))
			if err != nil {
				log.Println(err)
				return
			}
			m, err := MailCode(mail)
			if err != nil {
				log.Println(err)
				return
			}
			if err := SendMail(m); err != nil {
				log.Println(err)
				return
			}
		})
		if err := c.Reader.CommitMessages(context.Background(), m); err != nil {
			log.Println(err)
		}
	}
}

func (c *MailConsumer) Close() error {
	if err := c.Reader.Close(); err != nil {
		return err
	}
	return nil
}
