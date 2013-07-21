package graph

// Ungraph is an adjacency list undirected graph. It consumes 2E + V
// spaces.
type Ungraph struct {
	v   int
	e   *int
	adj [][]int
}

// NewGraph returns a Graph of size v implemented with an adjacency vertex
// list.
func NewGraph(v int) Ungraph {
	return Ungraph{
		v:   v,
		e:   new(int),
		adj: make([][]int, v),
	}
}

// AddEdge adds an edge from v to w. This is O(1).
func (a Ungraph) AddEdge(v, w int) {
	a.adj[v] = append(a.adj[v], w)
	a.adj[w] = append(a.adj[w], v)
	(*a.e)++
}

// Adj is a slice of vertices adjacent to v. This is O(E).
func (a Ungraph) Adj(v int) []int {
	return a.adj[v]
}

// V is the number of vertices. This is O(1).
func (a Ungraph) V() int {
	return a.v
}

// E is the number of edges. This is O(1).
func (a Ungraph) E() int {
	return *a.e
}

// GoString represents this graph as a string.
func (a Ungraph) GoString() string {
	return stringify(a)
}
