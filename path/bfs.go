package path

import (
	"container/list"
	"errors"
	"github.com/aybabtme/graph"
)

type bfs struct {
	g      graph.Graph
	from   int
	edgeTo []int
	marked []bool
}

// BuildBFS builds a Breath First Search PathFinder for graph g starting from
// source s
func BuildBFS(g graph.Graph, s int) (PathFinder, error) {

	var b bfs

	if s < 0 {
		return b, errors.New("Can't start BFS from negative source")
	}

	if s >= g.V() {
		return b, errors.New("Can't start BFS from vertex v >= total vertex count")
	}

	b = bfs{
		g:      g,
		from:   s,
		edgeTo: make([]int, g.V()),
		marked: make([]bool, g.V()),
	}

	queue := list.New()
	queue.PushBack(b.from)
	b.marked[b.from] = true

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

func (b bfs) HasPathTo(to int) bool {
	return b.marked[to]
}

func (b bfs) PathTo(to int) []int {
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
