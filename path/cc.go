package path

import (
	"github.com/aybabtme/graph"
)

type CC interface {
	// Connected tells whether vertices v and w are connected
	Connected(v, w int) bool
	// Count is the number of components in the graph
	Count() int
	// Id of the component containing vertex v
	Id(v int) int
}

type connectedComp struct {
	id    []int
	count int
}

func BuildCC(g graph.Graph) CC {

	cc := connectedComp{
		id:    make([]int, g.V()),
		count: 0,
	}

	marked := make([]bool, g.V())

	var visit func(v int)

	visit = func(v int) {
		marked[v] = true
		cc.id[v] = cc.count
		for _, adj := range g.Adj(v) {
			if !marked[adj] {
				visit(adj)
			}
		}
	}

	for v := 0; v < g.V(); v++ {
		if !marked[v] {
			visit(v)
			cc.count++
		}
	}

	return cc
}

func (c connectedComp) Connected(v, w int) bool {
	return c.id[v] == c.id[w]
}

func (c connectedComp) Count() int {
	return c.count
}

func (c connectedComp) Id(v int) int {
	return c.id[v]
}
