package path

import (
	"errors"
	"github.com/aybabtme/graph"
)

type tremauxDFS struct {
	g      graph.Graph
	from   int
	marked []bool
	edgeTo []int
}

// BuildDFS builds a depth first search PathFinder for graph g starting from
// source s
func BuildDFS(g graph.Graph, s int) (PathFinder, error) {

	var t tremauxDFS

	if s < 0 {
		return t, errors.New("Can't start DFS from negative source")
	}

	if s >= g.V() {
		return t, errors.New("Can't start DFS from vertex v >= total vertex count")
	}

	t = tremauxDFS{
		g:      g,
		from:   s,
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
	}

	var visit func(v int)

	visit = func(v int) {
		t.marked[v] = true
		for _, adj := range g.Adj(v) {
			if !t.marked[adj] {
				t.edgeTo[adj] = v
				visit(adj)
			}
		}
	}

	visit(t.from)

	return t, nil
}

func (t tremauxDFS) HasPathTo(to int) bool {
	return t.marked[to]
}

func (t tremauxDFS) PathTo(to int) []int {
	if !t.HasPathTo(to) {
		return []int{}
	}

	var path []int
	for next := to; next != t.from; next = t.edgeTo[next] {
		path = append(path, next)
	}
	path = append(path, t.from)

	reverse(path)

	return path
}
