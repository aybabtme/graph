package graph

type adjList struct {
	v   int
	e   int
	adj [][]int
}

func NewAdjListGraph(v int) Graph {
	return adjList{
		v:   v,
		e:   0,
		adj: make([][]int, v),
	}
}

func (a adjList) AddEdge(v, w int) {
	a.adj[v] = append(a.adj[v], w)
	a.adj[w] = append(a.adj[w], v)
	a.e++
}

func (a adjList) Adj(v int) []int {
	return a.adj[v]
}

func (a adjList) V() int {
	return a.v
}

func (a adjList) E() int {
	return a.e
}

func (a adjList) GoString() string {
	return stringify(a)
}

type adjMatrix struct {
	matrix [][]bool
}

func NewMatrixGraph(v int) Graph {

	matrix := make([][]bool, v)

	for i := 0; i < v; i++ {
		matrix[i] = make([]bool, v)
	}

	return adjMatrix{matrix: matrix}
}

func (a adjMatrix) AddEdge(v, w int) {
	a.matrix[v][w] = true
	a.matrix[w][v] = true
}

func (a adjMatrix) Adj(v int) []int {
	var vertices []int
	for i, ok := range a.matrix[v] {
		if ok {
			vertices = append(vertices, i)
		}
	}
	return vertices
}

func (a adjMatrix) V() int {
	return len(a.matrix)
}

func (a adjMatrix) E() int {
	c := 0
	for v := range a.matrix {
		c += len(a.Adj(v))
	}
	return c
}

func (a adjMatrix) GoString() string {
	return stringify(a)
}
