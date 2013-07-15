package path

import (
	"fmt"
	"github.com/aybabtme/graph"
)

type PathFinder interface {
	HasPathTo(to int) bool
	PathTo(to int) []int
}

type TremauxDFS struct {
	g      graph.Graph
	from   int
	marked []bool
	edgeTo []int
}

func BuildTremauxDFS(g graph.Graph, from int) PathFinder {

	t := TremauxDFS{
		g:      g,
		from:   from,
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
	}

	var visit func(v int)

	steps := 0

	visit = func(v int) {

		t.marked[v] = true
		for _, adj := range g.Adj(v) {
			steps++
			if !t.marked[adj] {
				t.edgeTo[adj] = v
				visit(adj)
			}
		}
	}

	visit(from)

	return t
}

func (t TremauxDFS) HasPathTo(to int) bool {
	return t.marked[to]
}

func (t TremauxDFS) PathTo(to int) []int {
	return []int{}
}
