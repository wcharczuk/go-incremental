package incr

import (
	"context"
)

// Map returns a new map incremental.
func Map2[A, B, C comparable](i0 Incr[A], i1 Incr[B], fn func(A, B) C) Incr[C] {
	m2 := &map2Incr[A, B, C]{
		i0: i0,
		i1: i1,
		fn: fn,
	}
	m2.n = NewNode(
		m2,
		OptNodeChildOf(i0),
		OptNodeChildOf(i1),
	)
	return m2
}

type map2Incr[A, B, C comparable] struct {
	n     *Node
	i0    Incr[A]
	i1    Incr[B]
	fn    func(A, B) C
	value C
}

func (m *map2Incr[A, B, C]) Value() C {
	return m.value
}

func (m *map2Incr[A, B, C]) Stabilize(ctx context.Context, g Generation) error {
	oldValue := m.value
	m.value = m.fn(m.i0.Value(), m.i1.Value())
	if oldValue != m.value {
		m.n.changedAt = g
	}
	return nil
}

func (m *map2Incr[A, B, C]) Node() *Node {
	return m.n
}
