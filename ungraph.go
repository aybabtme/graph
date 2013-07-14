package girafe

type AdjList struct {
	v   int
	adj [][]int
}

func NewAdjList(v int) Graph {
	return AdjList{
		v:   v,
		adj: make([][]int, v),
	}
}

func (a AdjList) AddEdge(v, w int) {
	a.adj[v] = append(a.adj[v], w)
	a.adj[w] = append(a.adj[w], v)
}

func (a AdjList) Adj(v int) []int {
	return a.adj[v]
}

func (a AdjList) V() int {
	return a.v
}
func (a AdjList) E() int {
	e := 0
	for _, v := range a.adj {
		e += len(v)
	}
	return e / 2
}

func (a AdjList) GoString() string {
	return stringify(a)
}

// AdjMatrix is a graph represented as a vertice-by-vertice matrix where the
// entries represent the existence or absence of edges
type AdjMatrix struct {
	matrix [][]bool
}

func NewAdjMatrix(v int) Graph {

	matrix := make([][]bool, v)

	for i := 0; i < v; i++ {
		matrix[i] = make([]bool, v)
	}

	return AdjMatrix{matrix: matrix}
}

func (a AdjMatrix) AddEdge(v, w int) {
	a.matrix[v][w] = true
	a.matrix[w][v] = true
}

func (a AdjMatrix) Adj(v int) []int {
	var edges []int
	for i, ok := range a.matrix[v] {
		if ok {
			edges = append(edges, i)
		}
	}
	return edges
}

func (a AdjMatrix) V() int {
	return len(a.matrix)
}

func (a AdjMatrix) E() int {
	c := 0
	for v := range a.matrix {
		c += len(a.Adj(v))
	}
	return c
}

func (a AdjMatrix) GoString() string {
	return stringify(a)
}
