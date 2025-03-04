package transformer

func Transform[T any](in <-chan T, action func(T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			out <- action(v)
		}
	}()

	return out
}
