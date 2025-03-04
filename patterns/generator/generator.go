package generator

func Generator[T any](values ...T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, v := range values {
			out <- v
		}
	}()

	return out
}
