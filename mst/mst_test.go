package mst

import (
	. "github.com/aybabtme/graph"
	"testing"
)

var (
	tinyEdgeWeightedGraph = []Edge{
		NewEdge(4, 5, .35),
		NewEdge(4, 7, .37),
		NewEdge(5, 7, .28),
		NewEdge(0, 7, .16),
		NewEdge(1, 5, .32),
		NewEdge(0, 4, .38),
		NewEdge(2, 3, .17),
		NewEdge(1, 7, .19),
		NewEdge(0, 2, .26),
		NewEdge(1, 2, .36),
		NewEdge(1, 3, .29),
		NewEdge(2, 7, .34),
		NewEdge(6, 2, .40),
		NewEdge(3, 6, .52),
		NewEdge(6, 0, .58),
		NewEdge(6, 4, .93),
	}
	// Kruskal and Prim will give the same tree, but the tree might be
	// ordered differently (children be left-right inverted)
	expectedKruskalMST = []Edge{
		NewEdge(0, 7, .16),
		NewEdge(2, 3, .17),
		NewEdge(1, 7, .19),
		NewEdge(0, 2, .26),
		NewEdge(5, 7, .28),
		NewEdge(4, 5, .35),
		NewEdge(6, 2, .40),
	}
	expectedPrimMST = []Edge{
		NewEdge(0, 7, .16),
		NewEdge(1, 7, .19),
		NewEdge(0, 2, .26),
		NewEdge(2, 3, .17),
		NewEdge(5, 7, .28),
		NewEdge(4, 5, .35),
		NewEdge(6, 2, .40),
	}
)

func TestKruskalMatchesKnownOutput(t *testing.T) {
	testMstAgainstKnownOutput(t, BuildKruskalMST, expectedKruskalMST)
}

func TestLazyPrimMatchesKnownOutput(t *testing.T) {
	testMstAgainstKnownOutput(t, BuildLazyPrimMST, expectedPrimMST)
}

func testMstAgainstKnownOutput(t *testing.T, mstFunc func(*WeightGraph) MST, expectedMST []Edge) {
	wg := NewWeightGraph(8)
	for _, edge := range tinyEdgeWeightedGraph {
		wg.AddEdge(edge)
	}

	mst := mstFunc(&wg)

	actualWeight := mst.Weight()

	if actualWeight != expectedWeight(expectedMST) {
		t.Errorf("Expected weight of MST to be %f but was %f",
			expectedWeight(expectedMST), actualWeight)
	}

	actualMST := mst.Edges()

	if len(actualMST) != len(expectedMST) {
		t.Fatalf("Expected MST to have length %d but was %d",
			len(expectedMST), len(actualMST))
	}

	for e := 0; e < len(expectedMST); e++ {
		expectedEdge := expectedMST[e]
		actualEdge := actualMST[e]

		if expectedEdge != actualEdge {
			t.Errorf("Expected to get edge %#v but was %#v", expectedEdge, actualEdge)
		}
	}
}

func expectedWeight(expectedMST []Edge) (w float64) {
	for _, edge := range expectedMST {
		w += edge.Weight()
	}
	return
}
