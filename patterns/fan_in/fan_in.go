package fan_in

import (
	"sync"
)

func MergeChannels[T any](channels ...<-chan T) chan T {
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	out := make(chan T)

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
