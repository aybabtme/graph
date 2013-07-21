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
