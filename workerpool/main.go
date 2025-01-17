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
	results := make(chan R)
	wg := new(sync.WaitGroup)

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
					case results <- transform(v):

					}
				}
			}
		}()
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	return results
}
