package path

import (
	"github.com/aybabtme/graph"
	"testing"
)

type pfFact func(graph.Graph, int) (PathFinder, error)

var pathFinders = []pfFact{
	BuildDFS,
	BuildBFS,
}

func TestSearchWithSimpleDisconnectedGraph(t *testing.T) {
	for _, pf := range pathFinders {
		simpleGraphHarness(t, pf)
	}

}

func simpleGraphHarness(t *testing.T, pfFactory pfFact) {
	g := graph.NewGraph(13)

	expectPathTo := []int{0, 1, 2, 3, 4, 5, 6}
	expectNoPathTo := []int{7, 8, 9, 10, 11, 12}

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 6)
	g.AddEdge(0, 5)
	g.AddEdge(6, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)

	g.AddEdge(7, 8)

	g.AddEdge(9, 10)
	g.AddEdge(9, 12)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)

	for _, v := range expectPathTo {
		pf, err := pfFactory(g, v)
		if err != nil {
			t.Fatalf("Couldn't build PathFinder, %v", err)
		}

		for _, conn := range expectPathTo {
			if !pf.HasPathTo(conn) {
				t.Errorf("Vertex %d should be connected to %d in graph %v",
					0, conn, g)
			}

			path := pf.PathTo(conn)

			l := len(path)
			if v == conn && l != 1 {
				t.Errorf("Path to itself should be [%d], path was %v", v, path)
			}
			if v != conn && l == 1 {
				t.Errorf("Path should be longer than 1 from %d to %d in graph %v",
					v, conn, g)
			}

			if l <= 0 {
				t.Errorf("Should have a len(path)>0 from %d to %d in graph %v",
					v, conn, g)
			}

			if l > len(expectPathTo) {
				t.Errorf("Path from %d to %d can have len()<=%d at most,"+
					" but path is %v",
					v, conn, len(expectPathTo), path)
			}
		}

		for _, disconn := range expectNoPathTo {
			if pf.HasPathTo(disconn) {
				t.Errorf("Vertex %d should NOT be connected to %d in graph %v",
					0, disconn, g)
			}

			path := pf.PathTo(disconn)

			if len(path) != 0 {
				t.Errorf("Should NOT have path from %d to %d, was %v",
					v, disconn, g)
			}
		}
	}
}

var shouldHaveErr = []struct {
	size   int
	source int
	msg    string
}{
	{0, -10, "Empty graph with negative index (-10)"},
	{0, -1, "Empty graph with negative index (-1)"},
	{0, 0, "Empty graph with index of 0"},
	{0, 1, "Empty graph with index too big (1)"},
	{0, 10, "Empty graph with index too big (10)"},
	{1, 1, "Graph size 1 with index too big (1)"},
	{1, 2, "Graph size 1 with index too big (2)"},
	{1, 10, "Graph size 1 with index too big (10)"},
	{1, -1, "Graph size 1 with negative index (-1)"},
	{1, -2, "Graph size 1 with negative index (-2)"},
	{1, -10, "Graph size 1 with negative index (-10)"},
	{10, 10, "Graph size 10 with index too big (10)"},
	{10, 11, "Graph size 10 with index too big (11)"},
	{10, 100, "Graph size 10 with index too big (100)"},
	{10, -1, "Graph size 10 with negative index (-1)"},
	{10, -2, "Graph size 10 with negative index (-2)"},
	{10, -10, "Graph size 10 with negative index (-10)"},
}

func TestSearchPanicBadArgsForSource(t *testing.T) {
	for _, pf := range pathFinders {
		for _, tt := range shouldHaveErr {
			badArgsHarness(t, pf, tt.size, tt.source, tt.msg)
		}
	}
}

func badArgsHarness(
	t *testing.T,
	pf pfFact,
	size, source int,
	msg string,
) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test \"%s\" panicked with message, %v", msg, r)
		}
	}()

	if _, err := pf(graph.NewGraph(size), source); err == nil {
		t.Errorf("%s, should have error", msg)
	}
}
