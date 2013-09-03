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

	var v int
	n, err := fmt.Fscanf(input, "%d\n", &v)
	if err != nil {
		return Ungraph{}, fmt.Errorf("Failed reading vertex count, %v", err)
	} else if n != 1 {
		return Ungraph{}, fmt.Errorf("Wanted to read %d integer from vertex count, read %d", 1, n)
	}

	g := NewGraph(v)

	var e int
	n, err = fmt.Fscanf(input, "%d\n", &e)
	if err != nil {
		return Ungraph{}, fmt.Errorf("Failed reading edge count, %v", err)
	} else if n != 1 {
		return Ungraph{}, fmt.Errorf("Wanted to read %d integer from edge count, read %d", 1, n)
	}

	readEdgePair := func(num int) (int, int, error) {
		var from, to int
		n, err := fmt.Fscanf(input, "%d %d\n", &from, &to)
		if err != nil {
			return -1, -1, fmt.Errorf("Failed reading edge line #%d, %v", num, err)
		} else if n != 2 {
			return -1, -1, fmt.Errorf("Wanted to read %d integers from edge line, read %d", 2, n)
		}
		return from, to, nil
	}

	for i := 0; i < e; i++ {
		from, to, err := readEdgePair(i)
		if err != nil {
			return g, err
		}
		g.AddEdge(from, to)
	}

	return g, nil
}
