package typed

import (
	"github.com/aybabtme/graph"
)

// WeightGraph is a graph with weighted edges.
type WeightGraph struct {
	cur    int
	index  map[interface{}]int
	invIdx []interface{}
	g      *graph.WeightGraph
}

// NewWeightGraph creates an empty graph with v vertices
func NewWeightGraph(v int) *WeightGraph {
	return &WeightGraph{
		cur:    0,
		index:  make(map[interface{}]int, v),
		invIdx: make([]interface{}, v),
		g:      graph.NewWeightGraph(v),
	}
}

// AddEdge adds weigthed edge e to this graph
func (wg *WeightGraph) AddEdge(e Edge) {
	v := e.Either()
	w := e.Other(v)
	vID, ok := wg.index[v]
	if !ok {
		vID = wg.cur
		wg.index[v] = vID
		wg.invIdx[vID] = v
		wg.cur++
	}

	wID, ok := wg.index[w]
	if !ok {
		wID = wg.cur
		wg.index[w] = wID
		wg.invIdx[wID] = w
		wg.cur++
	}

	wg.g.AddEdge(graph.NewEdge(vID, wID, e.Weight()))
}

// Adj gives the edges incident to v
func (wg *WeightGraph) Adj(v interface{}) []Edge {
	vID := wg.index[v]
	intEdges := wg.g.Adj(vID)
	edges := make([]Edge, len(intEdges))
	for i, e := range intEdges {
		vID := e.Either()
		wID := e.Other(vID)

		edges[i] = NewEdge(
			wg.invIdx[vID],
			wg.invIdx[wID],
			e.Weight(),
		)
	}
	return edges
}

// Edges gives all the edges in this graph
func (wg *WeightGraph) Edges() []Edge {
	intEdges := wg.g.Edges()
	edges := make([]Edge, len(intEdges))
	for i, e := range intEdges {
		vID := e.Either()
		wID := e.Other(vID)

		edges[i] = NewEdge(
			wg.invIdx[vID],
			wg.invIdx[wID],
			e.Weight(),
		)
	}
	return edges
}

// V is the number of vertives
func (wg *WeightGraph) V() int {
	return wg.g.V()
}

// E is the number of edges
func (wg *WeightGraph) E() int {
	return wg.g.E()
}

// Edge is a weighted edge in a weighted graph
type Edge struct {
	weight float64
	from   interface{}
	to     interface{}
}

// NewEdge creates a weigthed edge to be used by a WeightGraph
func NewEdge(v, w interface{}, weight float64) Edge {
	return Edge{weight: weight, from: v, to: w}
}

// Less tells if this edge is less than the other edge
func (e *Edge) Less(other Edge) bool {
	return e.weight < other.weight
}

// Either returns either vertices of this edge.
func (e *Edge) Either() interface{} {
	return e.from
}

// Other tells the other end of this edge, from v's perspective.
func (e *Edge) Other(v interface{}) interface{} {
	if e.from == v {
		return e.to
	}
	return e.from
}

// Weight tells the weight of this edge
func (e *Edge) Weight() float64 {
	return e.weight
}
