package mst

import (
	"container/heap"
	"github.com/aybabtme/graph"
)

type lazyPrim struct {
	tree   []graph.Edge
	weight float64
}

// BuildLazyPrimMST builds the MST for a weighted graph wg.
// This is O(E log E) and extra space proportional to E.
func BuildLazyPrimMST(wg *graph.WeightGraph) MST {
	var p lazyPrim
	var pq *edgePQ
	heap.Init(pq)
	marked := make([]bool, wg.V())

	var visit func(int)
	visit = func(v int) {
		marked[v] = true
		for _, adj := range wg.Adj(v) {
			if !marked[adj.Other(v)] {
				heap.Push(pq, adj)
			}
		}
	}

	for {
		if pq.Len() == 0 {
			return p
		}

		e := heap.Pop(pq).(graph.Edge)
		v := e.Either()
		w := e.Other(v)

		if marked[v] && marked[w] {
			continue
		}

		p.tree = append(p.tree, e)

		if !marked[v] {
			visit(v)
		}
		if !marked[w] {
			visit(w)
		}
	}
}

func (p lazyPrim) Edges() []graph.Edge {
	return p.tree
}

func (p lazyPrim) Weight() float64 {
	return p.weight
}
