package typed

import (
	"github.com/aybabtme/graph"
)

// Ungraph is an adjacency list undirected graph. It consumes 2E + V
// spaces.
type Ungraph struct {
	cur    int
	index  map[interface{}]int
	invIdx []interface{}
	g      graph.Ungraph
}

// NewGraph returns a Graph of size v implemented with an adjacency vertex
// list.
func NewGraph(v int) *Ungraph {
	return &Ungraph{
		cur:    0,
		index:  make(map[interface{}]int, v),
		invIdx: make([]interface{}, v),
		g:      graph.NewGraph(v),
	}
}

// AddEdge adds an edge from v to w. This is O(1).
func (un *Ungraph) AddEdge(v, w interface{}) {
	vID, ok := un.index[v]
	if !ok {
		vID = un.cur
		un.index[v] = vID
		un.invIdx[vID] = v
		un.cur++
	}

	wID, ok := un.index[w]
	if !ok {
		wID = un.cur
		un.index[w] = wID
		un.invIdx[wID] = w
		un.cur++
	}

	un.g.AddEdge(vID, wID)
}

// Adj is a slice of vertices adjacent to v. This is O(E).
func (un *Ungraph) Adj(v interface{}) []interface{} {
	// find ID
	vID := un.index[v]

	// delegate
	adjIdx := un.g.Adj(vID)

	// transform []idx to []interface{}
	adj := make([]interface{}, len(adjIdx))
	for i, idx := range adjIdx {
		adj[i] = un.invIdx[idx]
	}
	return adj
}

// ID returns the integer representation of the vertex
func (un *Ungraph) ID(v interface{}) int {
	return un.index[v]
}

// V is the number of vertices. This is O(1).
func (un *Ungraph) V() int {
	return un.g.V()
}

// E is the number of edges. This is O(1).
func (un *Ungraph) E() int {
	return un.g.E()
}

// GoString represents this graph as a string.
func (un *Ungraph) GoString() string {
	return stringify(un)
}
