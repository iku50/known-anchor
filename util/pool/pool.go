package pool

import "sync"

type Job func()
type Pool struct {
	JobQueue   chan Job
	wg         sync.WaitGroup
	closed     bool
	maxWorkers int
	mu         sync.Mutex
}

func NewPool(size int) *Pool {
	p := &Pool{
		JobQueue:   make(chan Job),
		maxWorkers: size,
	}
	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer p.wg.Done()
			for {
				select {
				case job, ok := <-p.JobQueue:
					if !ok {
						return
					}
					job()
				case <-p.shutdownSignal():
					return
				}
			}
		}()
	}
	return p
}

func (p *Pool) Submit(job Job) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.closed {
		p.JobQueue <- job
	}
}

func (p *Pool) Shutdown() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.closed {
		close(p.JobQueue)
		p.closed = true
	}
	p.wg.Wait()
}

func (p *Pool) shutdownSignal() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(ch)
	}()
	return ch
}
