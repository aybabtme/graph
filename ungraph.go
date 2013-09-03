package graph

import (
	"fmt"
	"io"
)

// Ungraph is an adjacency list undirected graph. It consumes 2E + V
// spaces.
type Ungraph struct {
	v   int
	e   *int
	adj [][]int
}

// AddEdge adds an edge from v to w. This is O(1).
func (a Ungraph) AddEdge(v, w int) {
	a.adj[v] = append(a.adj[v], w)
	a.adj[w] = append(a.adj[w], v)
	(*a.e)++
}

// Adj is a slice of vertices adjacent to v. This is O(E).
func (a Ungraph) Adj(v int) []int {
	return a.adj[v]
}

// V is the number of vertices. This is O(1).
func (a Ungraph) V() int {
	return a.v
}

// E is the number of edges. This is O(1).
func (a Ungraph) E() int {
	return *a.e
}

// GoString represents this graph as a string.
func (a Ungraph) GoString() string {
	return stringify(a)
}

// NewGraph returns a Graph of size v implemented with an adjacency vertex
// list.
func NewGraph(v int) Ungraph {
	return Ungraph{
		v:   v,
		e:   new(int),
		adj: make([][]int, v),
	}
}

// ReadGraph constructs an undirected graph from the io.Reader expecting
// to find data formed such as:
//   v
//   e
//   a b
//   c d
//   ...
//   y z
// where `v` is the vertex count, `e` the number of edges and `a`, `b`, `c`,
// `d`, ..., `y` and `z` are edges between `a` and `b`, `c` and `d`, ..., and
// `y` and `z` respectively.
func ReadGraph(input io.Reader) (Ungraph, error) {

	scan := newGraphScanner(input)

	v, err := scan.NextInt()
	if err != nil {
		return Ungraph{}, fmt.Errorf("Failed reading vertex count. %v", err)
	}

	g := NewGraph(v)

	e, err := scan.NextInt()
	if err != nil {
		return Ungraph{}, fmt.Errorf("Failed reading edge count. %v", err)
	}

	for i := 0; i < e; i++ {
		from, to, err := scan.NextEdge()
		if err != nil {
			return g, fmt.Errorf("Failed at edge line=%d. %v", i, err)
		}
		g.AddEdge(from, to)
	}

	return g, nil
}
