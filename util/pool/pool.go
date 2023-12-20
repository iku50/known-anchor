package pool

import "sync"

type Job func()
type Pool struct {
	JobQueue chan Job
	wg       sync.WaitGroup
}

func NewPool(size int) *Pool {
	p := Pool{
		JobQueue: make(chan Job),
	}
	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			for job := range p.JobQueue {
				job()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Submit(job Job) {
	p.JobQueue <- job
}

func (p *Pool) Shutdown() {
	close(p.JobQueue)
	p.wg.Wait()
}
