package mq

import (
	"context"
	"log"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Conn    *ampq.Connection
	Channel *ampq.Channel
	Queue   *ampq.Queue
	Tag     string
	Done    chan error
}

func NewProducer() *Producer {
	p := Producer{
		Tag:  "producer",
		Done: make(chan error),
	}
	var err error
	p.Conn, err = ampq.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer p.Conn.Close()
	p.Channel, err = p.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer p.Channel.Close()

	*p.Queue, err = p.Channel.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return &p
}

func (p *Producer) Produce(from string, funcName string, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	msg := Message{
		From:      from,
		Func:      funcName,
		Content:   content,
		TimeStamp: time.Now().Unix(),
	}
	data, err := msg.toBytes()
	failOnError(err, "Failed to marshal message")
	err = p.Channel.PublishWithContext(ctx,
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		ampq.Publishing{
			DeliveryMode: ampq.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", data)
}
