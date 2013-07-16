package graph

import (
	"bytes"
	"errors"
	"strconv"
)

type Digraph interface {
	Graph
	Reverse() Digraph
}

type diAdjList struct {
	v   int
	e   int
	adj [][]int
}

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

type Dag interface {
	Digraph
	Sort() []int
}

type dag struct {
	diAdjList
}

func NewDag(d Digraph) (Dag, error) {
	di, ok := d.(diAdjList)
	if !ok {
		panic("Not an adjacency list digraph")
	}
	if len(DirectedCycle(di)) == 0 {
		return dag{di}, nil
	} else {
		return dag{}, errors.New("Digraph has at least one cycle")
	}
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

	return revPostOrder
}

func DirectedCycle(di Digraph) []int {

	marked := make([]bool, di.V())
	edgeTo := make([]int, di.V())
	onStack := make([]bool, di.V())

	var cycle []int
	hasCycle := func() bool {
		return len(cycle) != 0
	}

	var visit func(v int)

	visit = func(v int) {
		onStack[v] = true
		marked[v] = true
		for _, adj := range di.Adj(v) {
			if hasCycle() {
				return
			} else if !marked[v] {
				edgeTo[adj] = v
				visit(adj)
			} else if onStack[adj] {
				for x := v; x != adj; x = edgeTo[x] {
					cycle = append(cycle, x)
				}
				cycle = append(cycle, adj)
				cycle = append(cycle, v)
			}
			onStack[v] = false
		}
	}

	for v := 0; v < di.V(); v++ {
		if !marked[v] {
			visit(v)
			if hasCycle() {
				return cycle
			}
		}
	}

	return []int{}
}
