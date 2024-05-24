package future

import "sync"

type Future[T any] interface {
	Get() (T, error)
}

type futureImpl[T any] struct {
	wg    sync.WaitGroup
	value T
	err   error
}

func (f *futureImpl[T]) Get() (T, error) {
	f.wg.Wait()

	return f.value, f.err
}

// New
// Create a new Future running fn in a go routine.
func New[T any](fn func() (T, error)) Future[T] {

	f := futureImpl[T]{}
	f.wg.Add(1)

	go func() {
		f.value, f.err = fn()
		f.wg.Done()
	}()

	return &f
}
