package path

import (
	"github.com/aybabtme/graph"
)

// CC holds a representation of the connected components in a graph
type CC interface {
	// Connected tells whether vertices v and w are connected
	Connected(v, w int) bool
	// Count is the number of components in the graph
	Count() int
	// ID of the component containing vertex v
	ID(v int) int
}

type connectedComp struct {
	id    []int
	count int
}

// BuildCC builds a Connected Component representation of graph g
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

func (c connectedComp) ID(v int) int {
	return c.id[v]
}

// SCC holds a representation of the strongly connected components in a
// directed graph
type SCC interface {
	// Connected tells whether vertices v and w are connected
	StronglyConnected(v, w int) bool
	// Count is the number of components in the digraph
	Count() int
	// ID of the component containing vertex v
	ID(v int) int
}

type strongComp struct {
	id    []int
	count int
}

// BuildSCC builds a Strongly Connected Component representation of digraph di
// using the Kosaraju Sharir algorithm.
func BuildSCC(di graph.Digraph) SCC {
	scc := strongComp{
		id:    make([]int, di.V()),
		count: 0,
	}

	marked := make([]bool, di.V())

	var dfs func(v int)

	dfs = func(v int) {
		marked[v] = true
		scc.id[v] = scc.count
		for _, adj := range di.Adj(v) {
			if !marked[adj] {
				dfs(adj)
			}
		}
	}

	dfo := BuildDFO(di)
	for _, v := range dfo.ReversePost {
		if !marked[v] {
			dfs(v)
			scc.count++
		}
	}

	return scc
}

func (scc strongComp) StronglyConnected(v, w int) bool {
	return scc.id[v] == scc.id[w]
}

func (scc strongComp) Count() int {
	return scc.count
}

func (scc strongComp) ID(v int) int {
	return scc.id[v]
}

// DFO holds information regarding the paths in a graph when traversed in
// depth-first search.
type DFO struct {
	// Pre is a slice of the vertices in preorder.
	Pre []int
	// Post is a slice of the vertices in postorder.
	Post []int
	// ReversePost is a slive of the vertices in reverse postorder.
	ReversePost []int
}

// BuildDFO constructs a depth first order representation of
// digraph di.
func BuildDFO(di graph.Digraph) *DFO {
	dfo := &DFO{}
	marked := make([]bool, di.V())

	var dfs func(int)

	dfs = func(v int) {
		dfo.Pre = append(dfo.Pre, v)

		marked[v] = true
		for _, adj := range di.Adj(v) {
			if !marked[adj] {
				dfs(adj)
			}
		}

		dfo.Post = append(dfo.Post, v)
		dfo.ReversePost = append(dfo.ReversePost, v)
	}

	for i := 0; i < di.V(); i++ {
		if !marked[i] {
			dfs(i)
		}
	}

	reverse(dfo.ReversePost)

	return dfo
}
