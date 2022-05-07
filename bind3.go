package incr

import (
	"context"
	"fmt"
)

// Bind3 lets you swap out an entire subgraph of a computation based
// on a given function and two inputs.
func Bind3[A, B, C, D any](a Incr[A], b Incr[B], c Incr[C], fn func(A, B, C) Incr[D]) Incr[D] {
	o := &bind3Incr[A, B, C, D]{
		n:  NewNode(),
		a:  a,
		b:  b,
		c:  c,
		fn: fn,
	}
	Link(o, a, b, c)
	return o
}

var (
	_ Incr[bool]   = (*bind3Incr[string, float64, uint64, bool])(nil)
	_ GraphNode    = (*bind3Incr[string, float64, uint64, string])(nil)
	_ Stabilizer   = (*bind3Incr[string, float64, uint64, string])(nil)
	_ fmt.Stringer = (*bind3Incr[string, float64, uint64, string])(nil)
)

type bind3Incr[A, B, C, D any] struct {
	n     *Node
	a     Incr[A]
	b     Incr[B]
	c     Incr[C]
	fn    func(A, B, C) Incr[D]
	bind  Incr[D]
	value D
}

func (b *bind3Incr[A, B, C, D]) Node() *Node { return b.n }

func (b *bind3Incr[A, B, C, D]) Value() D { return b.value }

func (b *bind3Incr[A, B, C, D]) SetBind(v Incr[D]) {
	b.bind = v
}

func (b *bind3Incr[A, B, C, D]) Bind() (oldValue, newValue Incr[D]) {
	return b.bind, b.fn(b.a.Value(), b.b.Value(), b.c.Value())
}

func (b *bind3Incr[A, B, C, D]) Stabilize(ctx context.Context) error {
	BindUpdate[D](ctx, b)
	b.value = b.bind.Value()
	return nil
}

func (b *bind3Incr[A, B, C, D]) String() string {
	return "bind3[" + b.n.id.Short() + "]"
}
