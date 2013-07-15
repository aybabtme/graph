package path

import (
	"github.com/aybabtme/graph"
	"testing"
)

func TestDFSWithSimpleDisconnectedGraph(t *testing.T) {
	g := graph.NewAdjList(13)

	expectPathTo := []int{1, 2, 3, 4, 5, 6}
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
		pf := BuildTremauxDFS(g, v)

		for _, conn := range expectPathTo {
			if !pf.HasPathTo(conn) {
				t.Errorf("Vertex %d should be connected to %d in graph %v",
					0, conn, g)
			}
		}

		for _, disconn := range expectNoPathTo {
			if pf.HasPathTo(disconn) {
				t.Errorf("Vertex %d should NOT be connected to %d in graph %v",
					0, disconn, g)
			}
		}
	}

}
