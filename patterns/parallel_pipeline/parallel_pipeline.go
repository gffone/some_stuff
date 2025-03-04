package parallel_pipeline

import "sync"

func generate[T any](values ...T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, v := range values {
			out <- v
		}
	}()

	return out
}

func process[T any](in chan T, action func(T) T, n int) chan T {
	out := make(chan T)
	wg := &sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- action(v)
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func ParallelPipeline[T any](action func(T) T, parallelFactor int, values ...T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range process(generate(values...), action, parallelFactor) {
			out <- v
		}
	}()

	return out
}
