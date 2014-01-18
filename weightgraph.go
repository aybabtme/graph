package graph

import (
	"bytes"
	"fmt"
	"io"
)

// WeightGraph is a graph with weighted edges.
type WeightGraph struct {
	adj [][]Edge
	e   int
}

// NewWeightGraph creates an empty graph with v vertices
func NewWeightGraph(v int) *WeightGraph {
	return &WeightGraph{
		adj: make([][]Edge, v),
		e:   0,
	}
}

// ReadWeightGraph constructs an undirected graph from the io.Reader expecting
// to find data formed such as:
//   v
//   e
//   a b w0
//   c d w1
//   ...
//   y z wN
// where `v` is the vertex count, `e` the number of edges and `a`, `b`, `c`,
// `d`, ..., `y` and `z` are edges between `a` and `b`, `c` and `d`, ..., and
// `y` and `z` respectively, and `wN` is the weight of that edge.
func ReadWeightGraph(input io.Reader) (*WeightGraph, error) {
	scan := newWeighGraphScanner(input)

	v, err := scan.NextInt()
	if err != nil {
		return &WeightGraph{}, fmt.Errorf("failed reading vertex count, %v", err)
	}

	g := NewWeightGraph(v)

	e, err := scan.NextInt()
	if err != nil {
		return &WeightGraph{}, fmt.Errorf("failed reading edge count, %v", err)
	}

	for i := 0; i < e; i++ {
		from, to, weight, err := scan.NextEdge()
		if err != nil {
			return g, fmt.Errorf("failed at edge line=%d, %v", i, err)
		}
		g.AddEdge(NewEdge(from, to, weight))

	}

	return g, nil
}

// AddEdge adds weigthed edge e to this graph
func (wg *WeightGraph) AddEdge(e Edge) {
	wg.adj[e.from] = append(wg.adj[e.from], e)
	wg.adj[e.to] = append(wg.adj[e.to], e)
	wg.e++
}

// Adj gives the edges incident to v
func (wg *WeightGraph) Adj(v int) []Edge {
	return wg.adj[v]
}

// Edges gives all the edges in this graph
func (wg *WeightGraph) Edges() []Edge {
	var edges []Edge
	for v := 0; v < wg.V(); v++ {
		for _, adj := range wg.Adj(v) {
			if adj.Other(v) > v {
				edges = append(edges, adj)
			}
		}
	}
	return edges
}

// V is the number of vertives
func (wg *WeightGraph) V() int {
	return len(wg.adj)
}

// E is the number of edges
func (wg *WeightGraph) E() int {
	return wg.e
}

// GoString represents this weighted graph
func (wg *WeightGraph) GoString() string {
	var output bytes.Buffer

	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	for v := 0; v < wg.V(); v++ {
		for _, w := range wg.Adj(v) {
			do(output.WriteString(w.GoString()))
			do(output.WriteRune('\n'))
		}
	}
	return output.String()
}

// Edge is a weighted edge in a weighted graph
type Edge struct {
	weight float64
	from   int
	to     int
}

// NewEdge creates a weigthed edge to be used by a WeightGraph
func NewEdge(v, w int, weight float64) Edge {
	return Edge{weight: weight, from: v, to: w}
}

// Less tells if this edge is less than the other edge
func (e *Edge) Less(other Edge) bool {
	return e.weight < other.weight
}

// Either returns either vertices of this edge.
func (e *Edge) Either() int {
	return e.from
}

// Other tells the other end of this edge, from v's perspective.
func (e *Edge) Other(v int) int {
	if e.from == v {
		return e.to
	}
	return e.from
}

// Weight tells the weight of this edge
func (e *Edge) Weight() float64 {
	return e.weight
}

// GoString represents this edge in a directed, weighted fashion
func (e *Edge) GoString() string {
	return fmt.Sprintf("%d-%d %.5f", e.from, e.to, e.weight)
}
