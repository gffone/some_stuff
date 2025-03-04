package barrier

import "sync"

type Barrier struct {
	mu    sync.Mutex
	count int
	size  int

	beforeCh chan struct{}
	afterCh  chan struct{}
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		size:     size,
		beforeCh: make(chan struct{}, size),
		afterCh:  make(chan struct{}, size),
	}
}

func (b *Barrier) Before() {
	b.mu.Lock()
	b.count++
	if b.count == b.size {
		for range b.size {
			b.beforeCh <- struct{}{}
		}
	}

	b.mu.Unlock()
	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mu.Lock()
	b.count--
	if b.count == b.size {
		for range b.size {
			b.afterCh <- struct{}{}
		}
	}

	b.mu.Unlock()
	<-b.afterCh
}
