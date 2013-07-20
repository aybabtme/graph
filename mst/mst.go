package mst

import (
	"github.com/aybabtme/graph"
)

// MST is the minimum spanning tree of a WeightedGraph.  If the edges have
// unique weights, the MST will be unique.  Otherwise, this MST is one of the
// MSTs for the graph.
type MST interface {
	// Edges in the MST
	Edges() []graph.Edge
	// Weight gives the total weight of the MST
	Weight() float64
}

type edgePQ []edgeItem

type edgeItem struct {
	edge  *graph.Edge
	index int
}

func (e edgePQ) Len() int {
	return len(e)
}

func (e edgePQ) Less(v, w int) bool {
	return e[v].edge.Less(*e[w].edge)
}

func (e edgePQ) Swap(v, w int) {
	e[v], e[w] = e[w], e[v]
	e[v].index = v
	e[w].index = w
}

func (e *edgePQ) Push(x interface{}) {
	n := e.Len()
	item := edgeItem{
		x.(*graph.Edge),
		n,
	}
	*e = append(*e, item)
}

func (e *edgePQ) Pop() interface{} {
	old := *e
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*e = old[0 : n-1]
	return item.edge
}
