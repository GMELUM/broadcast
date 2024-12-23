package pool

import (
	"sync"
	"time"
)

type Pool struct {
	task     chan func()
	sem      chan struct{}
	limiter  <-chan time.Time
	limitMux sync.Mutex
}

func New(size, spawn, limit uint32) *Pool {
	pool := &Pool{
		task:    make(chan func()),
		sem:     make(chan struct{}, size),
		limiter: time.NewTicker(time.Second / time.Duration(limit)).C,
	}

	for i := uint32(0); i < spawn; i++ {
		pool.sem <- struct{}{}
		go pool.worker(func() {})
	}

	return pool
}

func (pool *Pool) Schedule(task func()) {
	<-pool.limiter
	select {
	case pool.task <- task:
	case pool.sem <- struct{}{}:
		go pool.worker(task)
	}
}

func (pool *Pool) worker(task func()) {
	defer func() { <-pool.sem }()
	task()
	for task := range pool.task {
		task()
	}
}

// UpdateLimiter обновляет время лимитера
func (pool *Pool) UpdateLimiter(limit uint32) {
	pool.limitMux.Lock()
	defer pool.limitMux.Unlock()

	// Создаем новый тикер с новым лимитом
	pool.limiter = time.NewTicker(time.Second / time.Duration(limit)).C
}
