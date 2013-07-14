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

func degree(g Graph, v int) int {
	return len(g.Adj(v))
}

func maxDegree(g Graph) int {
	max := 0
	deg := 0
	for v := 0; v < g.V(); v++ {
		deg = degree(g, v)
		if deg > max {
			max = deg
		}
	}
	return max
}

func avgDegree(g Graph) float64 {
	e := float64(g.E())
	v := float64(g.V())
	return 2.0 * e / v
}

func numSelfLoop(g Graph) int {
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
