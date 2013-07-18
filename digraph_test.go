package graph

import (
	"testing"
)

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

func digraphWithoutCycle() Digraph {
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

func digraphWithCycle() Digraph {
	di := NewDigraph(4)

	di.AddEdge(0, 1)
	di.AddEdge(1, 3)
	di.AddEdge(0, 2)
	di.AddEdge(3, 2)
	di.AddEdge(2, 0)

	return di
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

func TestDigraphCanReverse(t *testing.T) {
	digraphWithCycle().Reverse()
	digraphWithoutCycle().Reverse()

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
