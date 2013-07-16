package graph

import (
	"bytes"
	"fmt"
	"strconv"
)

// Graph is a undirected graph with V vertices
type Graph interface {
	fmt.GoStringer
	// AddEdge adds an edge v-w
	AddEdge(v, w int)
	// Adj is a slice of vertices adjacent to v
	Adj(v int) []int
	// V is the number of vertices
	V() int
	// E is the number of edges
	E() int
}

func stringify(g Graph) string {
	var output bytes.Buffer

	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			do(output.WriteString(strconv.Itoa(v)))
			do(output.WriteRune('-'))
			do(output.WriteString(strconv.Itoa(w)))
			do(output.WriteRune('\n'))
		}
	}
	return output.String()
}

func Degree(g Graph, v int) int {
	return len(g.Adj(v))
}

func MaxDegree(g Graph) int {
	max := 0
	deg := 0
	for v := 0; v < g.V(); v++ {
		deg = Degree(g, v)
		if deg > max {
			max = deg
		}
	}
	return max
}

func MinDegree(g Graph) int {
	min := Degree(g, g.V())
	deg := 0
	for v := 0; v < g.V()-1; v++ {
		deg = Degree(g, v)
		if deg < min {
			min = deg
		}
	}
	return min
}

func AvgDegree(g Graph) float64 {
	e := float64(g.E())
	v := float64(g.V())
	return 2.0 * e / v
}

func NumSelfLoop(g Graph) int {
	c := 0
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v == w {
				c++
			}
		}
	}
	// Each edge counted twice
	return c / 2
}

func HasCycle(g Graph) bool {

	if NumSelfLoop(g) != 0 {
		return true
	}

	marked := make([]bool, g.V())
	var visit func(v, u int) bool

	visit = func(v, u int) bool {
		marked[v] = true
		for _, adj := range g.Adj(v) {
			if marked[adj] {
				visit(adj, v)
			} else if u != adj {
				return true
			}
		}
		return false
	}

	for v := 0; v < g.V(); v++ {
		if !marked[v] {
			if visit(v, v) {
				return true
			}
		}
	}
	return false
}

func IsBipartite(g Graph) bool {
	marked := make([]bool, g.V())
	color := make([]bool, g.V())

	var visit func(v int) bool

	visit = func(v int) bool {
		marked[v] = true
		for _, adj := range g.Adj(v) {
			if marked[adj] {
				color[v] = !color[adj]
				visit(adj)
			} else if color[v] == color[adj] {
				return false
			}
		}
		return true
	}

	for v := 0; v < g.V(); v++ {
		if !marked[v] {
			if !visit(v) {
				return false
			}
		}
	}

	return true
}
