package pipeline

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

func transform[T any](in chan T, action func(T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			out <- action(v)
		}
	}()

	return out
}

func Pipeline[T any](action func(T) T, values ...T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range transform(generate(values...), action) {
			out <- v
		}
	}()

	return out
}
