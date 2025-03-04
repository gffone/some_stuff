package tee

func Tee[T any](in <-chan T, n int) []chan T {
	outChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outChs[i] = make(chan T)
	}

	go func() {
		for val := range in {
			for _, ch := range outChs {
				ch <- val
			}
		}

		for _, ch := range outChs {
			close(ch)
		}
	}()

	return outChs
}
