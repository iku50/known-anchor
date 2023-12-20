package mq

import (
	"known-anchors/util/pool"
	"log"

	"known-anchors/util/mail"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Conn    *ampq.Connection
	Channel *ampq.Channel
	Tag     string
	Done    chan error
	Pool    *pool.Pool
}

// helper function to check amqp errors
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewConsumer() *Consumer {
	c := Consumer{
		Tag:  "consumer",
		Done: make(chan error),
	}
	var err error
	c.Conn, err = ampq.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer c.Conn.Close()
	c.Channel, err = c.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer c.Channel.Close()

	err = c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	c.Pool = pool.NewPool(5)
	return &c
}

func (c *Consumer) Consume() {
	msgs, err := c.Channel.Consume(
		"hello", // queue
		c.Tag,   // consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer")
	for {
		select {
		case msg := <-msgs:
			m, err := toMessage(msg.Body)
			log.Printf("Received a message: %s", msg.Body)
			if err != nil {
				log.Printf("Failed to unmarshal message: %s", err)
				msg.Ack(false)
				continue
			}
			switch m.Func {
			case "sendmail":
				c.Pool.Submit(func() {
					ma, err := mail.JsonToMail(m.Content)
					if err != nil {
						log.Printf("Failed to unmarshal message: %s", err)
						msg.Ack(false)
						return
					}
					if err := mail.SendMail(ma); err != nil {
						log.Printf("Failed to send mail: %s", err)
						msg.Ack(false)
						return
					}
				})
			default:
				msg.Ack(false)
			}

		case <-c.Done:
			return
		}
	}
}

func (c *Consumer) Shutdown() {
	c.Channel.Close()
	c.Conn.Close()
}
