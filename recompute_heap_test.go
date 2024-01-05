package incr

import (
	"testing"

	. "github.com/wcharczuk/go-incr/testutil"
)

func Test_recomputeHeap_Add(t *testing.T) {

	rh := newRecomputeHeap(32)

	n50 := newHeightIncr(5)
	n60 := newHeightIncr(6)
	n70 := newHeightIncr(7)

	rh.Add(n50)

	// assertions post add n50
	{
		ItsEqual(t, 1, rh.Len())
		ItsEqual(t, 1, rh.heights[5].Len())
		ItsNotNil(t, rh.heights[5].head)
		ItsNotNil(t, rh.heights[5].head.value)
		ItsEqual(t, n50.Node().id, rh.heights[5].head.value.Node().id)
		ItsEqual(t, true, rh.Has(n50))
		ItsEqual(t, false, rh.Has(n60))
		ItsEqual(t, false, rh.Has(n70))
		ItsEqual(t, 5, rh.MinHeight())
		ItsEqual(t, 5, rh.MaxHeight())
	}

	rh.Add(n60)

	// assertions post add n60
	{
		ItsEqual(t, 2, rh.Len())
		ItsEqual(t, 1, rh.heights[5].Len())
		ItsNotNil(t, rh.heights[5].head)
		ItsNotNil(t, rh.heights[5].head.value)
		ItsEqual(t, n50.Node().id, rh.heights[5].head.value.Node().id)
		ItsEqual(t, 1, rh.heights[6].Len())
		ItsNotNil(t, rh.heights[6].head)
		ItsNotNil(t, rh.heights[6].head.value)
		ItsEqual(t, n60.Node().id, rh.heights[6].head.value.Node().id)
		ItsEqual(t, true, rh.Has(n50))
		ItsEqual(t, true, rh.Has(n50))
		ItsEqual(t, false, rh.Has(n70))
		ItsEqual(t, 5, rh.MinHeight())
		ItsEqual(t, 6, rh.MaxHeight())
	}

	rh.Add(n70)

	// assertions post add n70
	{
		ItsEqual(t, 3, rh.Len())
		ItsEqual(t, 1, rh.heights[5].Len())
		ItsNotNil(t, rh.heights[5].head)
		ItsNotNil(t, rh.heights[5].head.value)
		ItsEqual(t, n50.Node().id, rh.heights[5].head.value.Node().id)
		ItsEqual(t, 1, rh.heights[6].Len())
		ItsNotNil(t, rh.heights[6].head)
		ItsNotNil(t, rh.heights[6].head.value)
		ItsEqual(t, n60.Node().id, rh.heights[6].head.value.Node().id)
		ItsEqual(t, 1, rh.heights[7].Len())
		ItsNotNil(t, rh.heights[7].head)
		ItsNotNil(t, rh.heights[7].head.value)
		ItsEqual(t, n70.Node().id, rh.heights[7].head.value.Node().id)
		ItsEqual(t, true, rh.Has(n50))
		ItsEqual(t, true, rh.Has(n60))
		ItsEqual(t, true, rh.Has(n70))
		ItsEqual(t, 5, rh.MinHeight())
		ItsEqual(t, 7, rh.MaxHeight())
	}
}

func Test_recomputeHeap_RemoveMin(t *testing.T) {
	rh := newRecomputeHeap(32)

	ItsNil(t, rh.RemoveMin())

	n00 := newHeightIncr(0)
	n01 := newHeightIncr(0)
	n10 := newHeightIncr(1)
	n100 := newHeightIncr(10)

	rh.Add(n00)
	ItsEqual(t, 1, rh.Len())
	ItsEqual(t, 32, len(rh.heights))
	ItsEqual(t, 1, rh.heights[0].Len())
	rh.Add(n01)
	ItsEqual(t, 2, rh.Len())
	ItsEqual(t, 32, len(rh.heights))
	ItsEqual(t, 2, rh.heights[0].Len())

	ItsEqual(t, 0, rh.minHeight)
	ItsEqual(t, 0, rh.maxHeight)

	rh.Add(n10)
	ItsEqual(t, 3, rh.Len())
	ItsEqual(t, 32, len(rh.heights))
	ItsEqual(t, 2, rh.heights[0].Len())
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 0, rh.minHeight)
	ItsEqual(t, 1, rh.maxHeight)

	rh.Add(n100)
	ItsEqual(t, 4, rh.Len())
	ItsEqual(t, 32, len(rh.heights))
	ItsEqual(t, 2, rh.heights[0].Len())
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 1, rh.heights[10].Len())
	ItsEqual(t, 0, rh.minHeight)
	ItsEqual(t, 10, rh.maxHeight)

	r00 := rh.RemoveMin()
	ItsNotNil(t, r00)
	ItsNotNil(t, r00.Node())
	ItsEqual(t, n00.n.id, r00.Node().id)
	ItsEqual(t, 3, rh.Len())
	ItsEqual(t, false, rh.Has(n00))
	ItsEqual(t, true, rh.Has(n01))
	ItsEqual(t, true, rh.Has(n10))
	ItsEqual(t, true, rh.Has(n100))
	ItsEqual(t, 1, rh.heights[0].Len())
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 1, rh.heights[10].Len())
	ItsEqual(t, 0, rh.minHeight)
	ItsEqual(t, 10, rh.maxHeight)

	r01 := rh.RemoveMin()
	ItsNotNil(t, r01)
	ItsNotNil(t, r01.Node())
	ItsEqual(t, n01.n.id, r01.Node().id)
	ItsEqual(t, 2, rh.Len())
	ItsEqual(t, false, rh.Has(n00))
	ItsEqual(t, false, rh.Has(n01))
	ItsEqual(t, true, rh.Has(n10))
	ItsEqual(t, true, rh.Has(n100))
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 1, rh.heights[10].Len())
	ItsEqual(t, 1, rh.minHeight)
	ItsEqual(t, 10, rh.maxHeight)

	r10 := rh.RemoveMin()
	ItsNotNil(t, r10)
	ItsNotNil(t, r10.Node())
	ItsEqual(t, n10.n.id, r10.Node().id)
	ItsEqual(t, 1, rh.Len())
	ItsEqual(t, false, rh.Has(n00))
	ItsEqual(t, false, rh.Has(n01))
	ItsEqual(t, false, rh.Has(n10))
	ItsEqual(t, true, rh.Has(n100))
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsEqual(t, 1, rh.heights[10].Len())
	ItsEqual(t, 10, rh.minHeight)
	ItsEqual(t, 10, rh.maxHeight)

	r100 := rh.RemoveMin()
	ItsNotNil(t, r100)
	ItsNotNil(t, r100.Node())
	ItsEqual(t, n100.n.id, r100.Node().id)
	ItsEqual(t, 0, rh.Len())
	ItsEqual(t, false, rh.Has(n00))
	ItsEqual(t, false, rh.Has(n01))
	ItsEqual(t, false, rh.Has(n10))
	ItsEqual(t, false, rh.Has(n100))
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsEqual(t, 0, rh.heights[10].Len())
	ItsEqual(t, 0, rh.minHeight)
	ItsEqual(t, 10, rh.maxHeight)
}

func Test_recomputeHeap_RemoveMinHeight(t *testing.T) {
	rh := newRecomputeHeap(10)

	n00 := newHeightIncr(0)
	n01 := newHeightIncr(0)
	n02 := newHeightIncr(0)

	n10 := newHeightIncr(1)
	n11 := newHeightIncr(1)
	n12 := newHeightIncr(1)
	n13 := newHeightIncr(1)

	n50 := newHeightIncr(5)
	n51 := newHeightIncr(5)
	n52 := newHeightIncr(5)
	n53 := newHeightIncr(5)
	n54 := newHeightIncr(5)

	rh.Add(n00)
	rh.Add(n01)
	rh.Add(n02)
	rh.Add(n10)
	rh.Add(n11)
	rh.Add(n12)
	rh.Add(n13)
	rh.Add(n50)
	rh.Add(n51)
	rh.Add(n52)
	rh.Add(n53)
	rh.Add(n54)

	ItsEqual(t, 12, rh.Len())
	ItsEqual(t, 0, rh.MinHeight())
	ItsEqual(t, 5, rh.MaxHeight())

	output := rh.RemoveMinHeight()
	ItsEqual(t, 9, rh.Len())
	ItsEqual(t, 3, len(output))
	ItsNil(t, rh.heights[0].head)
	ItsNil(t, rh.heights[0].tail)
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsNotNil(t, rh.heights[1].head)
	ItsNotNil(t, rh.heights[1].tail)
	ItsEqual(t, 4, rh.heights[1].Len())
	ItsNotNil(t, rh.heights[5].head)
	ItsNotNil(t, rh.heights[5].tail)
	ItsEqual(t, 5, rh.heights[5].Len())

	ItsEqual(t, 1, rh.MinHeight())
	ItsEqual(t, 5, rh.MaxHeight())

	output = rh.RemoveMinHeight()
	ItsEqual(t, 5, rh.Len())
	ItsEqual(t, 4, len(output))
	ItsNil(t, rh.heights[0].head)
	ItsNil(t, rh.heights[0].tail)
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsNil(t, rh.heights[1].head)
	ItsNil(t, rh.heights[1].tail)
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsNotNil(t, rh.heights[5].head)
	ItsNotNil(t, rh.heights[5].tail)
	ItsEqual(t, 5, rh.heights[5].Len())

	output = rh.RemoveMinHeight()
	ItsEqual(t, 0, rh.Len())
	ItsEqual(t, 5, len(output))
	ItsNil(t, rh.heights[0].head)
	ItsNil(t, rh.heights[0].tail)
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsNil(t, rh.heights[1].head)
	ItsNil(t, rh.heights[1].tail)
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsNil(t, rh.heights[5].head)
	ItsNil(t, rh.heights[5].tail)
	ItsEqual(t, 0, rh.heights[5].Len())

	rh.Add(n50)
	rh.Add(n51)
	rh.Add(n52)
	rh.Add(n53)
	rh.Add(n54)

	ItsEqual(t, 5, rh.Len())
	ItsEqual(t, 5, rh.MinHeight())
	ItsEqual(t, 5, rh.MaxHeight())

	output = rh.RemoveMinHeight()
	ItsEqual(t, 0, rh.Len())
	ItsEqual(t, 5, len(output))
	ItsNil(t, rh.heights[0].head)
	ItsNil(t, rh.heights[0].tail)
	ItsEqual(t, 0, rh.heights[0].Len())
	ItsNil(t, rh.heights[1].head)
	ItsNil(t, rh.heights[1].tail)
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsNil(t, rh.heights[5].head)
	ItsNil(t, rh.heights[5].tail)
	ItsEqual(t, 0, rh.heights[5].Len())
}

func Test_recomputeHeap_Remove(t *testing.T) {
	rh := newRecomputeHeap(10)
	n10 := newHeightIncr(1)
	n11 := newHeightIncr(1)
	n20 := newHeightIncr(2)
	n21 := newHeightIncr(2)
	n22 := newHeightIncr(2)
	n30 := newHeightIncr(3)

	// this should just return
	rh.Remove(n10)

	rh.Add(n10)
	rh.Add(n11)
	rh.Add(n20)
	rh.Add(n21)
	rh.Add(n22)
	rh.Add(n30)

	ItsEqual(t, true, rh.Has(n10))
	ItsEqual(t, true, rh.Has(n11))
	ItsEqual(t, true, rh.Has(n20))
	ItsEqual(t, true, rh.Has(n21))
	ItsEqual(t, true, rh.Has(n22))
	ItsEqual(t, true, rh.Has(n30))

	rh.Remove(n21)

	ItsEqual(t, 5, rh.Len())
	ItsEqual(t, true, rh.Has(n10))
	ItsEqual(t, true, rh.Has(n11))
	ItsEqual(t, true, rh.Has(n20))
	ItsEqual(t, false, rh.Has(n21))
	ItsEqual(t, true, rh.Has(n22))
	ItsEqual(t, true, rh.Has(n30))

	ItsEqual(t, 2, rh.heights[2].Len())
	ItsEqual(t, n20.Node().ID(), rh.heights[2].head.value.Node().ID())
	ItsEqual(t, n22.Node().ID(), rh.heights[2].tail.value.Node().ID())

	rh.Remove(n10)
	rh.Remove(n11)

	ItsEqual(t, 3, rh.Len())
	ItsEqual(t, 0, rh.heights[1].Len())
	ItsNil(t, rh.heights[1].head)
	ItsNil(t, rh.heights[1].tail)
	ItsEqual(t, 2, rh.MinHeight())
	ItsEqual(t, 3, rh.MaxHeight())
}

func Test_recomputeHeap_nextMinHeightUnsafe_noItems(t *testing.T) {
	rh := new(recomputeHeap)

	rh.minHeight = 1
	rh.maxHeight = 3

	next := rh.nextMinHeightUnsafe()
	ItsEqual(t, 0, next)
}

func Test_recomputeHeap_nextMinHeightUnsafe_pastMax(t *testing.T) {
	r0 := Return("hello")
	rh := newRecomputeHeap(4)
	rh.minHeight = 1
	rh.maxHeight = 3

	rh.lookup[r0.Node().id] = &listItem[Identifier, INode]{
		key:   r0.Node().id,
		value: r0,
	}
	next := rh.nextMinHeightUnsafe()
	ItsEqual(t, 0, next)
}

func Test_recomputeHeap_adjustHeights(t *testing.T) {
	rh := newRecomputeHeap(8)
	ItsEqual(t, 8, len(rh.heights))
	rh.adjustHeights(9) // we use (1) indexing!
	ItsEqual(t, 10, len(rh.heights))
}

func Test_recomputeHeap_addAdjustsHeights(t *testing.T) {
	rh := newRecomputeHeap(8)
	ItsEqual(t, 8, len(rh.heights))

	v0 := newHeightIncr(32)
	rh.Add(v0)
	ItsEqual(t, 33, len(rh.heights))
	ItsEqual(t, 32, rh.minHeight)
	ItsEqual(t, 32, rh.maxHeight)

	v1 := newHeightIncr(64)
	rh.Add(v1)
	ItsEqual(t, 65, len(rh.heights))
	ItsEqual(t, 32, rh.minHeight)
	ItsEqual(t, 64, rh.maxHeight)
}

func Test_recomuteHeap_Add_regression(t *testing.T) {
	// real world use case! it's insane!
	rh := newRecomputeHeap(8)

	v0 := newHeightIncr(1)
	v1 := newHeightIncr(1)
	m1 := newHeightIncr(0)
	o2 := newHeightIncr(3)

	rh.addUnsafe(m1)
	rh.addUnsafe(m1)
	rh.addUnsafe(o2)
	m1.n.height = 2
	rh.addUnsafe(m1)
	rh.addUnsafe(v0)
	rh.addUnsafe(v1)
	rh.addUnsafe(o2)

	ItsEqual(t, 1, rh.minHeight)

	var nodesInLists int
	for _, l := range rh.heights {
		if l != nil {
			nodesInLists += l.Len()
		}
	}
	ItsEqual(t, len(rh.lookup), nodesInLists)

	var seen []Identifier
	for len(rh.lookup) > 0 {
		n := rh.RemoveMin()
		ItsNotNil(t, n)
		seen = append(seen, n.Node().id)
	}
	ItsEqual(t, []Identifier{v0.n.id, v1.n.id, m1.n.id, o2.n.id}, seen)
}

func Test_recomputeHeap_Add_regression2(t *testing.T) {
	// another real world use case! also insane!
	rh := newRecomputeHeap(256)

	observer4945d288 := newHeightIncr(1)
	rh.Add(observer4945d288)
	observer87df48be := newHeightIncr(1)
	rh.Add(observer87df48be)
	mapf2cb6e46 := newHeightIncr(0)
	rh.Add(mapf2cb6e46)
	map26e9bfb2a := newHeightIncr(0)
	rh.Add(map26e9bfb2a)
	map26e9bfb2a.n.height = 2
	rh.Add(map26e9bfb2a)
	map2dfe7c676 := newHeightIncr(1)
	rh.Add(map2dfe7c676)
	map2aa9d55f9 := newHeightIncr(1)
	rh.Add(map2aa9d55f9)
	observerbaad6dd3 := newHeightIncr(1)
	rh.Add(observerbaad6dd3)
	map2aa3f9a14 := newHeightIncr(1)
	rh.Add(map2aa3f9a14)
	observer6e9e8864 := newHeightIncr(1)
	rh.Add(observer6e9e8864)
	varb35bfa8a := newHeightIncr(1)
	rh.Add(varb35bfa8a)
	var54b93408 := newHeightIncr(1)
	rh.Add(var54b93408)
	alwaysc83986c6 := newHeightIncr(0)
	rh.Add(alwaysc83986c6)
	cutoff9d454a57 := newHeightIncr(0)
	rh.Add(cutoff9d454a57)
	varc0898518 := newHeightIncr(1)
	rh.Add(varc0898518)
	cutoff9d454a57.n.height = 4
	rh.Add(cutoff9d454a57)
	mapf2cb6e46.n.height = 3
	rh.Add(mapf2cb6e46)
	alwaysc83986c6.n.height = 2
	rh.Add(alwaysc83986c6)
	map26e9bfb2a.n.height = 7
	rh.Add(map26e9bfb2a)
	map2aa3f9a14.n.height = 5
	rh.Add(map2aa3f9a14)
	observer6e9e8864.n.height = 6
	rh.Add(observer6e9e8864)
	map2dfe7c676.n.height = 6
	rh.Add(map2dfe7c676)
	map2aa9d55f9.n.height = 6
	rh.Add(map2aa9d55f9)
	observerbaad6dd3.n.height = 8
	rh.Add(observerbaad6dd3)

	// now start """stabilization"""

	ItsEqual(t, 1, rh.minHeight)
	ItsEqual(t, 5, rh.heights[1].Len())

	minHeightBlock := rh.RemoveMinHeight()
	ItsEqual(t, 5, len(minHeightBlock))
	ItsEqual(t, true, allHeight(minHeightBlock, 1))
}

func Test_recomputeHeap_fix(t *testing.T) {
	rh := newRecomputeHeap(8)
	v0 := newHeightIncr(2)
	rh.Add(v0)
	v1 := newHeightIncr(3)
	rh.Add(v1)
	v2 := newHeightIncr(4)
	rh.Add(v2)

	ItsEqual(t, 2, rh.minHeight)
	ItsEqual(t, 1, rh.heights[2].Len())
	ItsEqual(t, 1, rh.heights[3].Len())
	ItsEqual(t, 1, rh.heights[4].Len())
	ItsEqual(t, 4, rh.maxHeight)

	v0.n.height = 1
	rh.fix(v0.n.id)

	ItsEqual(t, 1, rh.minHeight)
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 0, rh.heights[2].Len())
	ItsEqual(t, 1, rh.heights[4].Len())
	ItsEqual(t, 4, rh.maxHeight)

	rh.fix(v0.n.id)
	ItsEqual(t, 1, rh.minHeight)
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 0, rh.heights[2].Len())
	ItsEqual(t, 1, rh.heights[4].Len())
	ItsEqual(t, 4, rh.maxHeight)

	v2.n.height = 5
	rh.fix(v2.n.id)

	ItsEqual(t, 1, rh.minHeight)
	ItsEqual(t, 1, rh.heights[1].Len())
	ItsEqual(t, 0, rh.heights[2].Len())
	ItsEqual(t, 1, rh.heights[5].Len())
	ItsEqual(t, 0, rh.heights[4].Len())
	ItsEqual(t, 5, rh.maxHeight)
}

func allHeight(values []INode, height int) bool {
	for _, v := range values {
		if v.Node().height != height {
			return false
		}
	}
	return true
}
