package graph

import (
	"bytes"
	"errors"
	"strconv"
)

// Digraph is a directed graph
type Digraph interface {
	Graph
	// Reverse returns the reverse of this digraph
	Reverse() Digraph
}

type diAdjList struct {
	v   int
	e   int
	adj [][]int
}

// NewDigraph returns a digraph with v vertices, all disconnected
func NewDigraph(v int) Digraph {
	return diAdjList{
		v:   v,
		e:   0,
		adj: make([][]int, v),
	}
}

func (di diAdjList) Reverse() Digraph {
	rev := NewDigraph(di.V())
	for v := 0; v < di.V(); v++ {
		for _, w := range di.Adj(v) {
			rev.AddEdge(w, v)
		}
	}
	return rev
}

// AddEdge adds an edge from v to w, but not from w to v
func (di diAdjList) AddEdge(v, w int) {
	di.adj[v] = append(di.adj[v], w)
	di.e++
}

func (di diAdjList) Adj(v int) []int {
	return di.adj[v]
}

func (di diAdjList) V() int {
	return di.v
}

func (di diAdjList) E() int {
	return di.e
}

func (di diAdjList) GoString() string {
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

// DAG is a directed acyclic graph
type DAG interface {
	Digraph
	Sort() []int
}

type dag struct {
	diAdjList
}

// NewDAG returns a DAG built from digraph d, if d has no cycle. Otherwise
// it returns an error
func NewDAG(d Digraph) (DAG, error) {
	di, _ := d.(diAdjList)
	if len(DirectedCycle(di)) == 0 {
		return dag{di}, nil
	}
	return dag{}, errors.New("Digraph has at least one cycle")

}

func (d dag) Sort() []int {
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

// DirectedCycle returns a cycle in digraph di, if there is one
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
