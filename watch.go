package incr

import "context"

// Watch returns a new watch incremental that tracks values for a given incremental.
func Watch[A any](i Incr[A]) *WatchIncr[A] {
	w := &WatchIncr[A]{
		incr: i,
	}
	return w
}

// WatchIncr is the implementation of `Watch`.
type WatchIncr[A any] struct {
	n      *Node
	incr   Incr[A]
	value  A
	values []A
}

// Value implements Incr[A].
func (w *WatchIncr[A]) Value() A {
	return w.value
}

// Stabilize implements Incr[A].
func (w *WatchIncr[A]) Stabilize(ctx context.Context) error {
	w.value = w.incr.Value()
	w.values = append(w.values, w.value)
	return nil
}

// Values returns the observed values.
func (w *WatchIncr[A]) Values() []A {
	return w.values
}

// Node implements Incr[A].
func (w *WatchIncr[A]) Node() *Node {
	return w.n
}

// String implements fmt.Stringer.
func (w *WatchIncr[A]) String() string { return "watch[" + w.n.id.Short() + "]" }
