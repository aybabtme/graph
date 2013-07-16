package graph

import (
	"bytes"
	"fmt"
	"strconv"
)

// Ungraph is a undirected graph with V vertices
type Ungraph interface {
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

func stringify(g Ungraph) string {
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

func Degree(g Ungraph, v int) int {
	return len(g.Adj(v))
}

func MaxDegree(g Ungraph) int {
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

func MinDegree(g Ungraph) int {
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

func AvgDegree(g Ungraph) float64 {
	e := float64(g.E())
	v := float64(g.V())
	return 2.0 * e / v
}

func NumSelfLoop(g Ungraph) int {
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

func IsBipartite(g Ungraph) bool {
	panic("not implemented")
	return false
}
