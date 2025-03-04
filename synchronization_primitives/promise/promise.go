package promise

type result[T any] struct {
	value T
	err   error
}

type Promise[T any] struct {
	resultCh chan result[T]
}

func NewPromise[T any](asyncFn func() (T, error)) Promise[T] {
	promise := Promise[T]{
		resultCh: make(chan result[T]),
	}

	go func() {
		defer close(promise.resultCh)

		val, err := asyncFn()
		promise.resultCh <- result[T]{val, err}
	}()

	return promise
}

func (p *Promise[T]) Then(successFn func(T), errorFn func(error)) {
	go func() {
		res := <-p.resultCh
		if res.err == nil {
			successFn(res.value)
		} else {
			errorFn(res.err)
		}
	}()
}
