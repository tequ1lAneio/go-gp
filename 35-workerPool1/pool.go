package workpool

import (
	"errors"
	"fmt"
	"sync"
)

const (
	defaultCapacity = 5
	maxCapacity     = 10
)

type Task func()

type Pool struct {
	capacity int

	active chan struct{}
	tasks  chan Task

	wg   sync.WaitGroup
	quit chan struct{}

	preAlloc bool
	block    bool
}

func New(capacity int, opts ...Option) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}),
	}

	for _, opt := range opts {
		opt(p)
	}

	fmt.Printf("workerpool start(preAlloc=%t)\n", p.preAlloc)

	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}
	}

	go p.run()
	return p
}

func (p *Pool) run() {
	idx := len(p.active)

	if !p.preAlloc {
	loop:
		for t := range p.tasks {
			p.returnTask(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}

		for {
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			}
		}
	}
}
func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("worker pool freed")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	default:
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrNoIdleWorkerInPool
	}
}

func (p *Pool) Free() {
	println(333)
}
