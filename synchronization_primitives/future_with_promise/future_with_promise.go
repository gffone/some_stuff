package future_with_promise

type Future[T any] struct {
	resultCh <-chan T
}

func NewFuture[T any](resultCh <-chan T) Future[T] {
	return Future[T]{
		resultCh: resultCh,
	}
}

func (f *Future[T]) Get() T {
	return <-f.resultCh
}

type Promise[T any] struct {
	resultCh chan T
}

func NewPromise[T any]() Promise[T] {
	return Promise[T]{
		resultCh: make(chan T),
	}
}

func (p *Promise[T]) Set(val T) {
	p.resultCh <- val
	close(p.resultCh)
}

func (p *Promise[T]) GetFuture() Future[T] {
	return NewFuture[T](p.resultCh)
}
