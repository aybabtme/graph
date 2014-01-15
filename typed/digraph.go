package typed

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/aybabtme/graph"
)

// Digraph is a directed graph implementation using an adjacency list
type Digraph struct {
	cur    int
	index  map[interface{}]int
	invIdx []interface{}
	g      graph.Digraph
}

// NewDigraph returns a digraph with v vertices, all disconnected
func NewDigraph(v int) *Digraph {
	return &Digraph{
		cur:    0,
		index:  make(map[interface{}]int, v),
		invIdx: make([]interface{}, v),
		g:      graph.NewDigraph(v),
	}
}

// AddEdge adds an edge from v to w, but not from w to v. This is O(1).
func (di *Digraph) AddEdge(v, w interface{}) {
	vID, ok := di.index[v]
	if !ok {
		vID = di.cur
		di.index[v] = vID
		di.invIdx[vID] = v
		di.cur++
	}

	wID, ok := di.index[w]
	if !ok {
		wID = di.cur
		di.index[w] = wID
		di.invIdx[wID] = w
		di.cur++
	}

	di.g.AddEdge(vID, wID)
}

// Adj is a slice of vertices adjacent to v. This is O(E)
func (di *Digraph) Adj(v interface{}) []interface{} {
	// find ID
	vID := di.index[v]

	// delegate
	adjIdx := di.g.Adj(vID)

	// transform []idx to []interface{}
	adj := make([]interface{}, len(adjIdx))
	for i, idx := range adjIdx {
		adj[i] = di.invIdx[idx]
	}
	return adj
}

// ID returns the integer representation of the vertex
func (di *Digraph) ID(v interface{}) int {
	return di.index[v]
}

// V is the number of vertices.
func (di *Digraph) V() int {
	return di.g.V()
}

// E is the number of edges.
func (di *Digraph) E() int {
	return di.g.E()
}

// Reverse returns the reverse of this digraph
func (di *Digraph) Reverse() *Digraph {
	rev := NewDigraph(di.V())
	for vID := 0; vID < di.V(); vID++ {
		v := di.invIdx[vID]
		for _, w := range di.Adj(v) {
			rev.AddEdge(w, v)
		}
	}
	return rev
}

// GoString represents this graph as a string.
func (di *Digraph) GoString() string {
	output := bytes.NewBuffer(nil)
	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	for v := 0; v < di.V(); v++ {
		for _, w := range di.Adj(v) {
			do(fmt.Fprintf(output, "%v->%v\n", v, w))
		}
	}
	return output.String()
}

// DAG is a directed acyclic graph implemented with an adjacency list
// digraph.
type DAG struct {
	*Digraph
	g graph.DAG
}

// NewDAG returns a DAG built from digraph d, if d has no cycle. Otherwise
// it returns an error.
func NewDAG(di *Digraph) (*DAG, error) {
	dag, err := graph.NewDAG(di.g)
	if err != nil {
		return nil, err
	}
	return &DAG{di, dag}, errors.New("digraph has at least one cycle")

}

// Sort gives the topological sort of this DAG.
func (d *DAG) Sort() []interface{} {
	vertices := d.g.Sort()

	values := make([]interface{}, len(vertices))
	for i, val := range vertices {
		values[i] = d.invIdx[val]
	}
	return values
}

// DirectedCycle returns a cycle in digraph di, if there is one.
func DirectedCycle(di *Digraph) []interface{} {

	cycle := graph.DirectedCycle(di.g)

	values := make([]interface{}, len(cycle))
	for i, val := range cycle {
		values[i] = di.invIdx[val]
	}
	return reverse(values)
}

func reverse(s []interface{}) []interface{} {
	for i := 0; i < len(s)/2; i++ {
		opposite := len(s) - 1 - i
		s[i], s[opposite] = s[opposite], s[i]
	}
	return s
}
