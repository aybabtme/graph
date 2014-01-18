package graph

import (
	"os"
	"sort"
	"testing"
)

// Assert that NewDigraph returns an implementation of Graph
var _ Graph = NewDigraph(0)

func TestDigraphHasVertices(t *testing.T) {
	expected := 3

	di := NewDigraph(expected)
	actual := di.V()
	if expected != actual {
		t.Fatalf("Expected %d vertices, got %d", expected, actual)
	}
}

func TestDigraphHasEdges(t *testing.T) {
	expected := 3

	di := NewDigraph(expected + 1)
	di.AddEdge(0, 1)
	di.AddEdge(1, 2)
	di.AddEdge(2, 3)

	actual := di.E()
	if expected != actual {
		t.Fatalf("Expected %d edges, got %d", expected, actual)
	}
}

func TestDigraphAddEdgeThenHasAdjacent(t *testing.T) {

	expected := []int{1, 2, 3, 4, 5}

	di := NewDigraph(len(expected) + 1)

	for _, to := range expected {
		di.AddEdge(0, to)
	}

	actual := di.Adj(0)
	sort.Sort(sort.IntSlice(actual))

	if len(expected) != len(actual) {
		t.Fatalf("Expected length of %d but was %d",
			len(expected), len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected vertice %d but was %d", expected[i], actual[i])
		}
	}
}

// This digraph is also a Dag:
//
//  +--+         +--+
//  |0 |-------->|1 |
//  +--+         +--+
//    |            |
//    v            v
//  +--+         +--+
//  |2 |<--------|3 |
//  +--+         +--+
var expectedOrder = []int{0, 1, 3, 2}

func digraphWithoutCycle() *Digraph {
	di := NewDigraph(4)

	di.AddEdge(0, 1)
	di.AddEdge(1, 3)
	di.AddEdge(0, 2)
	di.AddEdge(3, 2)

	return di
}

// This digraph is not a Dag:
//
//  +--+         +--+
//  |0 |-------->|1 |
//  +--+         +--+
//    ^            |
//    |            |
//    v            v
//  +--+         +--+
//  |2 |<--------|3 |
//  +--+         +--+
var expectedCycle = []int{2, 0, 1, 3, 2}

func digraphWithCycle() *Digraph {
	di := NewDigraph(4)

	di.AddEdge(0, 1)
	di.AddEdge(1, 3)
	di.AddEdge(0, 2)
	di.AddEdge(3, 2)
	di.AddEdge(2, 0)

	return di
}

func TestDigraphCanStringify(t *testing.T) {
	for _, f := range []func() *Digraph{
		digraphWithoutCycle,
		digraphWithCycle,
	} {
		g := f()
		g.GoString()
	}
}

func TestDigraphIsDag(t *testing.T) {
	_, err := NewDAG(digraphWithoutCycle())
	if err != nil {
		t.Errorf("Should be a DAG, %v", err)
	}
}

func TestDigraphIsNotDag(t *testing.T) {
	_, err := NewDAG(digraphWithCycle())
	if err == nil {
		t.Errorf("Shouldn't be a DAG, %v", err)
	}
}

func TestDagCanSort(t *testing.T) {
	dag, _ := NewDAG(digraphWithoutCycle())

	order := dag.Sort()

	if len(order) != len(expectedOrder) {
		t.Fatalf("Expected order len=%d but was %d",
			len(expectedOrder), len(order))
	}

	for i := 0; i < len(order); i++ {
		if expectedOrder[i] != order[i] {
			t.Fatalf("Expected %v but was %v", expectedOrder, order)
		}
	}

}

func TestDigraphWithCycleCanReverse(t *testing.T) {
	di := digraphWithCycle()
	di.Reverse()
}

func TestDigraphWithoutCycleCanReverse(t *testing.T) {
	di := digraphWithoutCycle()
	di.Reverse()
}

func TestDigraphHasNoCycle(t *testing.T) {
	cycle := DirectedCycle(digraphWithoutCycle())
	if len(cycle) != 0 {
		t.Errorf("Shouldn't have cycle but got one, %#v", cycle)
	}
}

func TestDigraphHasCycle(t *testing.T) {
	cycle := DirectedCycle(digraphWithCycle())
	if len(cycle) == 0 {
		t.Fatalf("Should have cycle but got none")
	}

	if len(cycle) != len(expectedCycle) {
		t.Errorf("Expected cycle len=%d but was %d",
			len(expectedCycle), len(cycle))
	}

	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	min := minInt(len(cycle), len(expectedCycle))
	for i := 0; i < min; i++ {
		if expectedCycle[i] != cycle[i] {
			t.Fatalf("Expected %v but was %v", expectedCycle, cycle)
		}
	}
}

func TestDigraphFromFile(t *testing.T) {
	filename := "data/tinyDG.txt"
	if testing.Short() {
		t.Skip("Not running test using " + filename)
	}
	fd, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed opening %s, %v", filename, err)
	}
	defer fd.Close()

	g, err := ReadDigraph(fd)
	if err != nil {
		t.Fatalf("Couldn't read graph from %s, %v", filename, err)
	}

	wantV := 13
	gotV := g.V()

	if wantV != gotV {
		t.Errorf("Vertex count, want %d got %d", wantV, gotV)
	}

	wantE := 22
	gotE := g.E()

	if wantE != gotE {
		t.Errorf("Edge count, want %d got %d", wantE, gotE)
	}

	for _, v := range g.Adj(6) {
		switch v {
		case 0:
			continue
		case 4:
			continue
		case 8:
			continue
		case 9:
			continue
		default:
			t.Errorf("6 should not be adjacent to %d", v)
		}
	}

	// v=1 has no out-edges
	for _, v := range g.Adj(1) {
		t.Errorf("1 should not be adjacent to %d", v)
	}
}
