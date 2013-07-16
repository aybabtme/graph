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
	adj [][]int
}

func NewDigraph(v int) Digraph {
	return diAdjList{
		v:   v,
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
}

func (di diAdjList) Adj(v int) []int {
	return di.adj[v]
}

func (di diAdjList) V() int {
	return di.v
}

func (di diAdjList) E() int {
	e := 0
	for _, v := range di.adj {
		e += len(v)
	}
	return e / 2
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
	if IsDag(di) {
		return dag{di}, nil
	} else {
		return dag{}, errors.New("Digraph has cycles")
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

func IsDag(di Digraph) bool {
	panic("not implemented")
	return false
}
