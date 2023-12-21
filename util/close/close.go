package close

import "fmt"

type Closeable interface {
	Close() error
}

type Close struct {
	S []Closeable
}

func NewClose() *Close {
	return &Close{
		S: []Closeable{},
	}
}

func (c *Close) AddCloseable(closeable Closeable) {
	fmt.Println("adding closeable")
	c.S = append(c.S, closeable)
}

func (c *Close) CloseAll() error {
	for _, closeable := range c.S {
		if err := closeable.Close(); err != nil {
			return err
		}
	}
	return nil
}
