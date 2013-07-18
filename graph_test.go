package graph

import (
	"strconv"
	"testing"
)

const (
	BipartiteUngraph int = iota
	BipartiteDigraph
	NotBipartiteUngraph
	NotBipartiteDigraph
	CycleGraph
	NotCycleGraph
)

var graphs = map[int][]struct {
	from int
	to   int
}{
	// This is a bipartite graph
	// +-+      +-+     +-+
	// |0+------+1+-----+2|
	// +-+      +-+     +-+
	// It doesn't have cycles
	BipartiteUngraph: {{0, 1}, {1, 2}},
	CycleGraph:       {{0, 1}, {1, 2}},

	// This is a bipartite digraph
	// +-+        +-+
	// |0|------->|1|
	// +-+        +-+
	//  ^          |
	//  |          |
	//  |          v
	// +-+        +-+
	// |3|<-------|2|
	// +-+        +-+
	BipartiteDigraph: {
		{0, 1}, {1, 2},
		{3, 0}, {2, 3},
	},

	// This is NOT a bipartite graph
	// +-+      +-+
	// |0+------+1|
	// +++      +++
	//  |        |
	//  |   +-+  |
	//  +---+2+--+
	//      +-+
	// It has a cycle
	NotBipartiteUngraph: {
		{0, 1},
		{1, 2},
		{2, 0},
	},
	NotCycleGraph: {
		{0, 1},
		{1, 2},
		{2, 0},
	},

	// This is NOT a bipartite digraph
	// +-+      +-+
	// |0|----->|1|
	// +-+      +-+
	//  ^        |
	//  |   +-+  |
	//  +---|2|+-+
	//      +-+
	NotBipartiteDigraph: {
		{0, 1},
		{1, 2},
		{2, 0},
	},
}

func buildGraph(graphBuilder func(int) Graph, name int) Graph {
	edgeList, ok := graphs[name]
	if !ok {
		panic("Not a graph name: " + strconv.Itoa(name))
	}

	g := graphBuilder(len(edgeList) + 1)

	for _, test := range edgeList {
		g.AddEdge(test.from, test.to)
	}
	return g
}

func buildDigraph(graphBuilder func(int) Digraph, name int) Digraph {
	edgeList, ok := graphs[name]
	if !ok {
		panic("Not a graph name: " + strconv.Itoa(name))
	}

	g := graphBuilder(len(edgeList) + 1)

	for _, test := range edgeList {
		g.AddEdge(test.from, test.to)
	}
	return g
}

func TestMatrixGraphIsBipartite(t *testing.T) {
	g := buildGraph(NewMatrixGraph, BipartiteUngraph)
	if !IsBipartite(g) {
		t.Fatalf("Matrix Graph should be bipartite, %#v", g)
	}
}

func TestAdjGraphIsBipartite(t *testing.T) {
	g := buildGraph(NewAdjListGraph, BipartiteUngraph)
	if !IsBipartite(g) {
		t.Fatalf("Adj Graph should be bipartite, %#v", g)
	}
}

func TestDigraphIsBipartite(t *testing.T) {
	di := buildDigraph(NewDigraph, BipartiteDigraph)
	if !IsBipartite(di) {
		t.Fatalf("Digraph should be bipartite, %#v", di)
	}
}

func TestMatrixGraphIsNotBipartite(t *testing.T) {
	g := buildGraph(NewMatrixGraph, NotBipartiteUngraph)
	if !IsBipartite(g) {
		t.Fatalf("Matrix Graph should be bipartite, %#v", g)
	}
}

func TestAdjGraphIsNotBipartite(t *testing.T) {
	g := buildGraph(NewAdjListGraph, NotBipartiteUngraph)
	if !IsBipartite(g) {
		t.Fatalf("Adj Graph should be bipartite, %#v", g)
	}
}

func TestDigraphIsNotBipartite(t *testing.T) {
	di := buildDigraph(NewDigraph, NotBipartiteDigraph)
	if IsBipartite(di) {
		t.Fatalf("Digraph should not be bipartite, %#v", di)
	}
}

func TestMatrixGraphHasCycle(t *testing.T) {
	g := buildGraph(NewMatrixGraph, CycleGraph)
	if HasCycle(g) {
		t.Fatalf("Matrix Graph should have cycle, %#v", g)
	}
}

func TestAdjGraphHasCycle(t *testing.T) {
	g := buildGraph(NewAdjListGraph, CycleGraph)
	if HasCycle(g) {
		t.Fatalf("Adj Graph should have cycle, %#v", g)
	}
}

func TestMatrixGraphHasNoCycle(t *testing.T) {
	g := buildGraph(NewMatrixGraph, NotCycleGraph)
	if !HasCycle(g) {
		t.Fatalf("Matrix Graph should not have cycle, %#v", g)
	}
}

func TestAdjGraphHasNoCycle(t *testing.T) {
	g := buildGraph(NewAdjListGraph, NotCycleGraph)
	if !HasCycle(g) {
		t.Fatalf("Adj Graph should not have cycle, %#v", g)
	}
}