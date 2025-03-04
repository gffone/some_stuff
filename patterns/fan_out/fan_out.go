package fan_out

func SplitChannel[T any](in <-chan T, n int) []chan T {
	outChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outChs[i] = make(chan T)
	}

	go func() {
		idx := 0
		for v := range in {
			outChs[idx] <- v
			idx = (idx + 1) % n
		}

		for _, ch := range outChs {
			close(ch)
		}
	}()

	return outChs
}
