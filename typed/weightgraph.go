package typed

import (
	"bytes"
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
func NewWeightGraph(v int) WeightGraph {
	return WeightGraph{
		cur:    0,
		index:  make(map[interface{}]int, v),
		invIdx: make([]interface{}, v),
		g:      graph.NewWeightGraph(v),
	}
}

// AddEdge adds weigthed edge e to this graph
func (wg *WeightGraph) AddEdge(e graph.Edge) {
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

	wg.g.AddEdge(e)
}

// Adj gives the edges incident to v
func (wg *WeightGraph) Adj(v interface{}) []graph.Edge {
	vID := wg.index[v]
	return wg.g.Adj(vID)
}

// Edges gives all the edges in this graph
func (wg *WeightGraph) Edges() []graph.Edge {
	return wg.g.Edges()
}

// V is the number of vertives
func (wg *WeightGraph) V() int {
	return wg.g.V()
}

// E is the number of edges
func (wg *WeightGraph) E() int {
	return wg.g.E()
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
