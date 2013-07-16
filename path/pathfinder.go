package path

import (
	"github.com/aybabtme/graph"
)

type PathFinder interface {
	HasPathTo(to int) bool
	PathTo(to int) []int
}

type tremauxDFS struct {
	g      graph.Graph
	from   int
	marked []bool
	edgeTo []int
}

func BuildDFS(g graph.Graph, from int) PathFinder {

	if from < 0 {
		panic("Can't start DFS from vertex v < 0")
	}

	if from >= g.V() {
		panic("Can't start DFS from vertex v >= total vertex count")
	}

	t := tremauxDFS{
		g:      g,
		from:   from,
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

	visit(from)

	return t
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

func reverse(s []int) {
	var opposite int
	for i := 0; i < len(s)/2; i++ {
		opposite = len(s) - 1 - i
		s[i], s[opposite] = s[opposite], s[i]
	}
}
