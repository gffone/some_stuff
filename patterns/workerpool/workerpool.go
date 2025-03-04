package workerpool

import (
	"context"
	"sync"
)

func Start[T, R any](
	ctx context.Context,
	workersCount int,
	input <-chan T,
	transform func(elem T) R,
) <-chan R {
	out := make(chan R)
	wg := &sync.WaitGroup{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-input:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- transform(v):
					}
				}
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
