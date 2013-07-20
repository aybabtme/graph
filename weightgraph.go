package graph

import (
	"bytes"
	"fmt"
)

// WeightGraph is a graph with weighted edges.
type WeightGraph struct {
	fmt.GoStringer
	adj [][]Edge
	e   int
}

// NewWeightGraph creates an empty graph with v vertices
func NewWeightGraph(v int) WeightGraph {
	return WeightGraph{
		adj: make([][]Edge, v),
		e:   0,
	}
}

// AddEdge adds weigthed edge e to this graph
func (w *WeightGraph) AddEdge(e Edge) {
	w.adj[e.from] = append(w.adj[e.from], e)
	w.adj[e.to] = append(w.adj[e.to], e)
	w.e++
}

// Adj gives the edges incident to v
func (w *WeightGraph) Adj(v int) []Edge {
	return w.adj[v]
}

// Edges gives all the edges in this graph
func (w *WeightGraph) Edges() []Edge {
	var edges []Edge
	for v := 0; v < w.V(); v++ {
		for _, adj := range w.Adj(v) {
			if adj.Other(v) > v {
				edges = append(edges, adj)
			}
		}
	}
	return edges
}

// V is the number of vertives
func (w *WeightGraph) V() int {
	return len(w.adj)
}

// E is the number of edges
func (w *WeightGraph) E() int {
	return w.e
}

// GoString represents this weighted graph
func (w *WeightGraph) GoString() string {
	var output bytes.Buffer

	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	for v := 0; v < w.V(); v++ {
		for _, w := range w.Adj(v) {
			do(output.WriteString(w.GoString()))
			do(output.WriteRune('\n'))
		}
	}
	return output.String()
}

// Edge is a weighted edge in a weighted graph
type Edge struct {
	fmt.GoStringer
	weight float64
	from   int
	to     int
}

// NewEdge constructs a new edge.
func NewEdge(v, w int, weight float64) Edge {
	return Edge{weight: weight, from: v, to: w}
}

// Less tells if this edge is less than the other edge
func (e *Edge) Less(other Edge) bool { return e.weight < other.weight }

// Either returns either vertices of this edge.
func (e *Edge) Either() int { return e.from }

// Other tells the other end of this edge, from v's perspective.
func (e *Edge) Other(v int) int {
	if e.from == v {
		return e.to
	}
	return e.from
}

// GoString represents this edge in a directed, weighted fashion
func (e *Edge) GoString() string {
	return fmt.Sprintf("%d-%d %.5f", e.from, e.to, e.weight)
}
