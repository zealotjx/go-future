package future

import "testing"

func TestFuture(t *testing.T) {
	f := New[int64](func() (int64, error) {
		return 10, nil
	})
	r, err := f.Get()
	t.Logf("result: %v", r)
	t.Logf("error: %v", err)
}
