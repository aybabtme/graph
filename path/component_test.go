package path

import (
	"github.com/aybabtme/graph"
	"testing"
)

func TestConnectedComponent(t *testing.T) {

	g := graph.NewGraph(13)

	componentA := []int{0, 1, 2, 3, 4, 5, 6}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 6)
	g.AddEdge(0, 5)
	g.AddEdge(6, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)

	componentB := []int{7, 8}
	g.AddEdge(7, 8)

	componentC := []int{9, 10, 11, 12}
	g.AddEdge(9, 10)
	g.AddEdge(9, 12)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)

	cc := BuildCC(g)

	if cc.Count() != 3 {
		t.Errorf("Should have 3 components, was %d", cc.Count())
	}

	checkAllConnected(t, cc, componentA)
	checkAllConnected(t, cc, componentB)
	checkAllConnected(t, cc, componentC)

	checkNoneConnected(t, cc, componentA, componentB)
	checkNoneConnected(t, cc, componentA, componentC)

	checkNoneConnected(t, cc, componentB, componentA)
	checkNoneConnected(t, cc, componentB, componentC)

	checkNoneConnected(t, cc, componentC, componentA)
	checkNoneConnected(t, cc, componentC, componentB)

}

func checkAllConnected(t *testing.T, cc CC, comp []int) {
	for _, first := range comp {
		for _, second := range comp {
			if !cc.Connected(first, second) {
				t.Errorf("Vertices %d and %d should be connected",
					first, second)
			}

			if cc.ID(first) != cc.ID(second) {
				t.Errorf("Vertices %d and %d should have same id",
					first, second)
			}
		}
	}
}

func checkNoneConnected(t *testing.T, cc CC, firstComp, secondComp []int) {
	for _, first := range firstComp {
		for _, second := range secondComp {
			if cc.Connected(first, second) {
				t.Errorf("Vertices %d and %d should NOT be connected",
					first, second)
			}

			if cc.ID(first) == cc.ID(second) {
				t.Errorf("Vertices %d and %d should NOT have same id",
					first, second)
			}
		}
	}
}

var (
	// We use this graph for DFO testing
	//
	//  /---------------------\
	//  |                     v
	// +--+   +--+   +--+    +--+      +--+   +--+
	// | 1|<--| 0|<--| 2|  /-| 6|<-----| 7|<--| 8|
	// +--+   +--+   +--+  | +--+      +--+   +--+
	//         |      |    |  \------\
	//         v      v    |         v
	//        +--+   +--+  |        +--+   +--+
	//        | 5|<--| 3|  |   /----| 9|-->|10|
	//        +--+   +--+  |   |    +--+   +--+
	//         |           |   |     |
	//         v           |   v     v
	//        +--+         |  +--+  +--+
	//        | 4|<--------/  |11|->|12|
	//        +--+            +--+  +--+
	//
	dfoGraphEdges = []struct{ from, to int }{
		// The order matters to match the expectations
		{0, 5}, {0, 1},
		{1, 6},
		{2, 0}, {2, 3},
		{3, 5},
		// 4 has no outgoing edges
		{5, 4},
		{6, 4}, {6, 9},
		{7, 6},
		{8, 7},
		{9, 11}, {9, 12}, {9, 10},
		{11, 12},
	}
	// We know its orderings should be :
	preOrder     = []int{0, 5, 4, 1, 6, 9, 11, 12, 10, 2, 3, 7, 8}
	postOrder    = []int{4, 5, 12, 11, 10, 9, 6, 1, 0, 3, 2, 7, 8}
	revPosrOrder = []int{8, 7, 2, 3, 0, 1, 6, 9, 10, 11, 12, 5, 4}
)

func TestDFOMatchesKnownOutput(t *testing.T) {
	di := graph.NewDigraph(13)
	for _, edge := range dfoGraphEdges {
		di.AddEdge(edge.from, edge.to)
	}

	dfo := BuildDFO(di)
	compareIntSlices(t, preOrder, dfo.Pre, "Preorder should match.")
	compareIntSlices(t, postOrder, dfo.Post, "Postorder should match.")
	compareIntSlices(t, revPosrOrder, dfo.ReversePost, "Reversed postorder should match.")
}

func compareIntSlices(t *testing.T, expected, actual []int, failMSg string) {
	if len(expected) != len(actual) {
		t.Errorf("%s Different length, expected %d but was %d",
			failMSg, len(expected), len(actual))
		return
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("%s slice[%d] mismatch, expected %d but was %d."+
				"\nExpected=%#v\nActual=%#v",
				failMSg, i, expected[i], actual[i], expected, actual)
			return
		}
	}
}

var (
	// We use this crazy digraph, that's also a DAG, for our SCC testing
	//
	//                         +------------------------------------------------------+
	//                         |           +------------------------------------------|
	//                         +-----------|---------------------------+              |
	//   +---------------------|----------+|                           |              |
	//   |         +----------+|          ||               +--------------------------|
	//   v         v          |v          |v               v           |              |
	// +--+      +--+<-+--+  +--+  +--+  +--+      +--+  +--+  +--+  +--+      +--+<-+--+     +--+
	// | 1|      | 2|  | 3|<-| 4|<-+ 5|<-| 0|      |10|<-| 9|<-+12|<-|11|      | 8|  | 6|<---+| 7|
	// +--+      +--+->+--+  +--+  +--+  +--+      +--+  +--+  +--+  +--+      +--+->+--+     +--+
	//             |     |          ^     ^          |    ^|    ^     ^                        |
	//             |     +----------+     |          +----||----+     |                        |
	//             +----------------------+               |+----------+                        |
	//                                                    +------------------------------------+
	// It's easy to see that is has 5 strongly connected components
	sccGraphEdges = []struct{ from, to int }{
		{0, 1}, {0, 5},
		// 1 is not connected to anything
		{2, 3}, {2, 0},
		{3, 2}, {3, 5},
		{4, 3}, {4, 2},
		{5, 4},
		{6, 8}, {6, 9}, {6, 0}, {6, 4},
		{7, 6}, {7, 9},
		{8, 6},
		{9, 10}, {9, 11},
		{10, 12},
		{11, 12}, {11, 4},
		{12, 9},
	}
	// We know the following are strongly connected components
	sccExpected = []struct {
		id   int
		comp []int
	}{
		{id: 0, comp: []int{1}},
		{id: 1, comp: []int{0, 2, 3, 4, 5}},
		{id: 2, comp: []int{9, 10, 11, 12}},
		{id: 3, comp: []int{6, 8}},
		{id: 4, comp: []int{7}},
	}
)

func TestSCCMatchesKnownOutput(t *testing.T) {
	di := graph.NewDigraph(13)
	for _, edge := range sccGraphEdges {
		di.AddEdge(edge.from, edge.to)
	}
	scc := BuildSCC(di)

	expectedCount := len(sccExpected)
	actualCount := scc.Count()

	if expectedCount != actualCount {
		t.Fatalf("Didn't find the expected count, wanted %d got %d",
			expectedCount, actualCount)
	}

	for i := 0; i < scc.Count(); i++ {
		expectedID := sccExpected[i].id
		for _, v := range sccExpected[i].comp {
			actualID := scc.ID(v)
			if expectedID != actualID {
				t.Errorf("Wanted ID=%d but was %d", expectedID, actualID)
			}
			for _, w := range sccExpected[i].comp {
				if !scc.StronglyConnected(v, w) {
					t.Errorf("Vertices %d and %d should be strongly connected",
						v, w)
				}
			}
		}
	}
}
