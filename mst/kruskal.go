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

// BuildKruskalMST builds the MST for a weighted graph wg.
// This is O(E log E). If edges arrive sorted, E log V.
func BuildKruskalMST(wg *graph.WeightGraph) MST {

	var k kruskal

	var pq edgePQ
	heap.Init(&pq)
	for _, e := range wg.Edges() {
		w := e
		heap.Push(&pq, &w)
	}

	uf := unionfind.BuildUF(wg.V())

	for {
		if pq.Len() == 0 || len(k.tree) > wg.V()-1 {
			break
		}

		e := heap.Pop(&pq).(*graph.Edge)
		v := e.Either()
		w := e.Other(v)

		if !uf.Connected(v, w) {
			uf.Union(v, w)
			k.tree = append(k.tree, *e)
			k.weight += e.Weight()
		}
	}

	return k
}

// Edges is the edges that form the MST of this weighted graph
func (k kruskal) Edges() []graph.Edge {
	return k.tree
}

func (k kruskal) Weight() float64 {
	return k.weight
}
