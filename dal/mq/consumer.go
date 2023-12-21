package mq

type Consumer interface {
	Consume()
	Close() error
}

// here is an example of how to use this interface
// func main() {
// 	c := NewConsumer()
// 	go c.Consume()
// 	defer c.Shutdown()
// }
