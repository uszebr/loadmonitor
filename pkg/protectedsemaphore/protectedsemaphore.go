package protectedsemaphore

import (
	"sync"
)

type ProtectedSemaphore struct {
	mu            sync.Mutex
	poolSize      poolSize
	semaphoreChan chan struct{}
}

type poolSize int

func newPoolSize(poolSizeInt int) poolSize {
	if poolSizeInt < 0 {
		return 0
	}
	return poolSize(poolSizeInt)
}

func New(initialPoolsize int) *ProtectedSemaphore {
	poolS := newPoolSize(initialPoolsize)
	return &ProtectedSemaphore{
		mu:            sync.Mutex{},
		poolSize:      poolS,
		semaphoreChan: make(chan struct{}, int(poolS)),
	}
}

func (ps *ProtectedSemaphore) Increase() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.semaphoreChan <- struct{}{}
}

func (ps *ProtectedSemaphore) Decrease() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	<-ps.semaphoreChan
}

func (ps *ProtectedSemaphore) SetSize(size int) {
	poolS := newPoolSize(size)
	ps.mu.Lock()
	defer ps.mu.Unlock()
	oldChan := ps.semaphoreChan
	ps.semaphoreChan = make(chan struct{}, int(poolS))
	ps.poolSize = poolS
	close(oldChan)
	//drain data from old channel if present
	for range oldChan {
	}
}

func (ps *ProtectedSemaphore) Close() {
	close(ps.semaphoreChan)
}

func (ps *ProtectedSemaphore) Drain() {
	for range ps.semaphoreChan {
	}
}
