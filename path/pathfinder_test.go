package path

import (
	"github.com/aybabtme/graph"
	"testing"
)

func TestDFSWithSimpleDisconnectedGraph(t *testing.T) {
	g := graph.NewAdjList(13)

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
		pf := BuildDFS(g, v)

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
				t.Errorf("Path should be longer than 1 from %d to %d in graph %v")
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
