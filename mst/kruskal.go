package mst

import (
	"container/heap"
	"github.com/aybabtme/graph"
	"github.com/aybabtme/graph/unionfind"
)

type kruskal struct {
	tree   []graph.Edge
	weight float64
}

// BuildKruskalMST builds the MST for weighted graph wg
func BuildKruskalMST(wg graph.WeightGraph) MST {

	var k kruskal

	var pq edgePQ
	heap.Init(pq)
	for _, e := range wg.Edges() {
		heap.Push(pq, e)
	}

	uf := unionfind.BuildUF(wg.V())

	for {
		if pq.Len() == 0 || len(k.tree) < wg.V()-1 {
			break
		}

		e := heap.Pop(pq).(graph.Edge)
		v := e.Either()
		w := e.Other(v)

		if uf.Connected(v, w) {
			uf.Union(v, w)
			k.tree = append(k.tree, e)
		}
	}

	return k
}

func (k kruskal) Edges() []graph.Edge {
	return k.tree
}

func (k kruskal) Weight() float64 {
	return k.weight
}
