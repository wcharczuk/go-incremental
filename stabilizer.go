package incr

import "context"

// Stabilizer is a type that can be stabilized.
type Stabilizer interface {
	Stabilize(context.Context, Generation) error
	Node() *Node
}
