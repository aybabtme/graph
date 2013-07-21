package graph

import (
	"testing"
)

const (
	bipartiteUngraph int = iota
	notBipartiteUngraph
	cycleUngraph
	notCycleUngraph
)

type edgeList []struct{ from, to int }

var graphs = map[int]edgeList{
	// This is a bipartite graph
	// +-+      +-+     +-+
	// |0+------+1+-----+2|
	// +-+      +-+     +-+
	// It doesn't have cycles
	bipartiteUngraph: {{0, 1}, {1, 2}},
	notCycleUngraph:  {{0, 1}, {1, 2}},

	// This is NOT a bipartite graph
	// +-+      +-+
	// |0+------+1|
	// +++      +++
	//  |        |
	//  |   +-+  |
	//  +---+2+--+
	//      +-+
	// It has a cycle
	notBipartiteUngraph: {
		{0, 1},
		{1, 2},
		{2, 0},
	},
	cycleUngraph: {
		{0, 1},
		{1, 2},
		{2, 0},
	},
}

func TestGraphIsBipartite(t *testing.T) {
	edgeList := graphs[bipartiteUngraph]

	g := NewGraph(len(edgeList) + 1)

	for _, edge := range edgeList {
		g.AddEdge(edge.from, edge.to)
	}

	if !IsBipartite(g) {
		t.Fatalf(" Graph should be bipartite, %#v", g)
	}
}

func TestGraphIsNotBipartite(t *testing.T) {
	edgeList := graphs[notBipartiteUngraph]
	g := NewGraph(len(edgeList) + 1)

	for _, edge := range edgeList {
		g.AddEdge(edge.from, edge.to)
	}
	if IsBipartite(g) {
		t.Fatalf(" Graph should be bipartite, %#v", g)
	}
}

func TestGraphHasCycle(t *testing.T) {
	edgeList := graphs[cycleUngraph]
	g := NewGraph(len(edgeList) + 1)

	for _, edge := range edgeList {
		g.AddEdge(edge.from, edge.to)
	}
	if !HasCycle(g) {
		t.Fatalf(" Graph should have cycle, %#v", g)
	}
}

func TestGraphHasNoCycle(t *testing.T) {
	edgeList := graphs[notCycleUngraph]
	g := NewGraph(len(edgeList) + 1)

	for _, edge := range edgeList {
		g.AddEdge(edge.from, edge.to)
	}
	if HasCycle(g) {
		t.Fatalf(" Graph should not have cycle, %#v", g)
	}
}

func iterStep(current int) int {
	if current == 0 {
		return 1
	}
	return current * 7
}

func TestDegreeOfListGraph(t *testing.T) {
	for i := 2; i < maxGraphSize; i = iterStep(i) {
		g := NewGraph(i)
		expected := i - 1
		for j := 1; j < i; j++ {
			g.AddEdge(0, j)
		}
		actual := Degree(g, 0)
		if actual != expected {
			t.Errorf("Expected degree %d but was %d", expected, actual)
		}
	}
}

func TestMaxDegreeOfListGraph(t *testing.T) {
	for i := 2; i < maxGraphSize; i = iterStep(i) {
		g := NewGraph(i)
		expected := i - 1
		for j := 1; j < i; j++ {
			g.AddEdge(0, j)
		}
		actual := MaxDegree(g)
		if actual != expected {
			t.Errorf("Expected max degree %d but was %d", expected, actual)
		}
	}
}

func TestMinDegreeOfListGraph(t *testing.T) {
	for i := 3; i < maxGraphSize; i = iterStep(i) {
		g := NewGraph(i)

		expected := 0
		for j := 1; j < i-1; /* <-- Skip one */ j++ {
			g.AddEdge(0, j)
		}

		actual := MinDegree(g)
		if actual != expected {
			t.Errorf("Expected min degree %d but was %d", expected, actual)
		}

		expected = 2
		for j := 0; j < i; j++ {
			g.AddEdge(g.V()-1, j)
		}

		actual = MinDegree(g)
		if actual != expected {
			t.Errorf("Expected min degree %d but was %d for size %d",
				expected, actual, i)
		}

	}
}

func TestAvgDegreeOfListGraph(t *testing.T) {
	for i := 2; i < maxGraphSize; i = iterStep(i) {
		g := NewGraph(i)

		proportion := i / 2
		expected := 0.0
		for v := 0; v < proportion; v++ {
			for w := 0; w < proportion; w++ {
				if v == w {
					continue
				}
				g.AddEdge(v, w)
				expected += 2.0
			}
		}
		expected /= float64(g.V())
		actual := AvgDegree(g)
		if actual != expected {
			t.Fatalf("Expected avg degree %f but was %f",
				expected, actual)
		}

	}
}
