package path

import (
	"container/list"
	"errors"
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

func BuildDFS(g graph.Graph, from int) (PathFinder, error) {

	var t tremauxDFS

	if from < 0 {
		return t, errors.New("Can't start DFS from negative source")
	}

	if from >= g.V() {
		return t, errors.New("Can't start DFS from vertex v >= total vertex count")
	}

	t = tremauxDFS{
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

func reverse(s []int) {
	var opposite int
	for i := 0; i < len(s)/2; i++ {
		opposite = len(s) - 1 - i
		s[i], s[opposite] = s[opposite], s[i]
	}
}

type BFS struct {
	g      graph.Graph
	from   int
	edgeTo []int
	marked []bool
}

func BuildBFS(g graph.Graph, from int) (PathFinder, error) {

	var b BFS

	if from < 0 {
		return b, errors.New("Can't start BFS from negative source")
	}

	if from >= g.V() {
		return b, errors.New("Can't start BFS from vertex v >= total vertex count")
	}

	b = BFS{
		g:      g,
		from:   from,
		edgeTo: make([]int, g.V()),
		marked: make([]bool, g.V()),
	}

	queue := list.New()
	queue.PushBack(from)
	b.marked[from] = true

	for el := queue.Front(); queue.Len() != 0; el = queue.Front() {
		v, ok := queue.Remove(el).(int)
		if !ok {
			panic("Failed to assert type int")
		}

		for _, adj := range g.Adj(v) {
			if !b.marked[adj] {
				queue.PushBack(adj)
				b.edgeTo[adj] = v
				b.marked[adj] = true
			}
		}
	}

	return b, nil
}

func (b BFS) HasPathTo(to int) bool {
	return b.marked[to]
}

func (b BFS) PathTo(to int) []int {
	if !b.HasPathTo(to) {
		return []int{}
	}

	var path []int
	for next := to; next != b.from; next = b.edgeTo[next] {
		path = append(path, next)
	}
	path = append(path, b.from)

	reverse(path)

	return path
}
