package graph

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Digraph is a directed graph implementation using an adjacency list
type Digraph struct {
	v   int
	e   *int
	adj [][]int
}

// NewDigraph returns a digraph with v vertices, all disconnected
func NewDigraph(v int) Digraph {
	return Digraph{
		v:   v,
		e:   new(int),
		adj: make([][]int, v),
	}
}

// ReadDigraph constructs an undirected graph from the io.Reader expecting
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
func ReadDigraph(input io.Reader) (Digraph, error) {

	var v int
	n, err := fmt.Fscanf(input, "%d\n", &v)
	if err != nil {
		return Digraph{}, fmt.Errorf("Failed reading vertex count, %v", err)
	} else if n != 1 {
		return Digraph{}, fmt.Errorf("Wanted to read %d integer from vertex count, read %d", 1, n)
	}

	g := NewDigraph(v)

	var e int
	n, err = fmt.Fscanf(input, "%d\n", &e)
	if err != nil {
		return Digraph{}, fmt.Errorf("Failed reading edge count, %v", err)
	} else if n != 1 {
		return Digraph{}, fmt.Errorf("Wanted to read %d integer from edge count, read %d", 1, n)
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

// AddEdge adds an edge from v to w, but not from w to v. This is O(1).
func (di Digraph) AddEdge(v, w int) {
	di.adj[v] = append(di.adj[v], w)
	(*di.e)++
}

// Adj is a slice of vertices adjacent to v. This is O(E)
func (di Digraph) Adj(v int) []int {
	return di.adj[v]
}

// V is the number of vertices.
func (di Digraph) V() int {
	return di.v
}

// E is the number of edges.
func (di Digraph) E() int {
	return *di.e
}

// Reverse returns the reverse of this digraph
func (di Digraph) Reverse() Digraph {
	rev := NewDigraph(di.V())
	for v := 0; v < di.V(); v++ {
		for _, w := range di.Adj(v) {
			rev.AddEdge(w, v)
		}
	}
	return rev
}

// GoString represents this graph as a string.
func (di Digraph) GoString() string {
	var output bytes.Buffer

	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	for v := 0; v < di.V(); v++ {
		for _, w := range di.Adj(v) {
			do(output.WriteString(strconv.Itoa(v)))
			do(output.WriteString("->"))
			do(output.WriteString(strconv.Itoa(w)))
			do(output.WriteRune('\n'))
		}
	}
	return output.String()
}

// DAG is a directed acyclic graph implemented with an adjacency list
// digraph.
type DAG struct {
	*Digraph
}

// NewDAG returns a DAG built from digraph d, if d has no cycle. Otherwise
// it returns an error.
func NewDAG(d Digraph) (DAG, error) {
	if len(DirectedCycle(d)) == 0 {
		return DAG{&d}, nil
	}
	return DAG{}, errors.New("Digraph has at least one cycle")

}

// Sort gives the topological sort of this DAG.
func (d *DAG) Sort() []int {
	// Why aren't we using the path.BuildDFO(di) function?
	// - It would result in an import cycle, that is why.

	marked := make([]bool, d.V())

	var revPostOrder []int

	var visit func(int)

	visit = func(v int) {
		marked[v] = true
		for _, adj := range d.Adj(v) {
			if !marked[adj] {
				visit(adj)
			}
		}
		revPostOrder = append(revPostOrder, v)
	}

	for v := 0; v < d.V(); v++ {
		if !marked[v] {
			visit(v)
		}
	}

	return reverse(revPostOrder)
}

// DirectedCycle returns a cycle in digraph di, if there is one.
func DirectedCycle(di Digraph) []int {

	marked := make([]bool, di.V())
	edgeTo := make([]int, di.V())
	onStack := make([]bool, di.V())
	var cycle []int
	hasCycle := func() bool {
		return len(cycle) != 0
	}

	var dfs func(v int)

	dfs = func(v int) {
		onStack[v] = true
		marked[v] = true
		for _, w := range di.Adj(v) {
			if hasCycle() {
				return
			} else if !marked[w] {
				edgeTo[w] = v
				dfs(w)
			} else if onStack[w] {
				for x := v; x != w; x = edgeTo[x] {
					cycle = append(cycle, x)
				}
				cycle = append(cycle, w)
				cycle = append(cycle, v)
			}
		}
		onStack[v] = false
	}

	for v := 0; v < di.V(); v++ {
		if !marked[v] {
			dfs(v)
		}
	}

	return reverse(cycle)
}

func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		opposite := len(s) - 1 - i
		s[i], s[opposite] = s[opposite], s[i]
	}
	return s
}
