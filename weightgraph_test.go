package graph

import (
	"math"
	"os"
	"testing"
)

func TestCanCreateWeightGraph(t *testing.T) {
	for v := 1; v < 100; v++ {
		wg := NewWeightGraph(v)
		wg.AddEdge(NewEdge(0, v-1, float64(v)))
		wg.GoString()
		if wg.V() != v {
			t.Errorf("Expected wg to have %d vertices bu had %d", v, wg.V())
		}
	}
}

func TestWeightGraphCanAddAndCountEdges(t *testing.T) {
	wg := NewWeightGraph(4)
	wg.AddEdge(NewEdge(0, 1, 0.1))
	wg.AddEdge(NewEdge(0, 2, 0.2))
	wg.AddEdge(NewEdge(0, 3, 0.3))
	wg.AddEdge(NewEdge(2, 3, 0.4))

	if wg.E() != 4 {
		t.Errorf("Expected %d edges but wg.E()=%d", 3, wg.E())
	}

	if len(wg.Adj(0)) != 3 {
		t.Errorf("Expected 3 adjacent to 0, but was %v", wg.Adj(0))
	}

	if len(wg.Edges()) != 4 {
		t.Errorf("Expected 4 edges, but was %v", wg.Edges())
	}
}

func TestEdgeHaveValues(t *testing.T) {
	// Small steps by 0.05% increment
	for i := -1000.0; i < 1000.0; i += 0.0005*math.Abs(i) + 0.0001 {
		edge := NewEdge(0, 1, i)
		if i != edge.Weight() {
			t.Errorf("Should have weight %f but was %f", i, edge.Weight())
		}
		v := edge.Either()
		w := edge.Other(v)

		if v != 0 {
			if v != 1 {
				t.Errorf("v should be 1")
			}
			if w != 0 {
				t.Errorf("w should be 0")
			}
		} else {
			if v != 0 {
				t.Errorf("v should be 0")
			}
			if w != 1 {
				t.Errorf("w should be 1")
			}
		}
	}
}

func TestEdgeCanCompare(t *testing.T) {
	small := NewEdge(0, 1, 0.1)
	medium := NewEdge(0, 2, 0.2)
	big := NewEdge(0, 3, 0.3)

	if !small.Less(medium) {
		t.Errorf("'%#v' should be less than '%#v' but was not ",
			small, medium)
	}

	if !small.Less(big) {
		t.Errorf("'%#v' should be less than '%#v' but was not ",
			small, big)
	}

	if !medium.Less(big) {
		t.Errorf("'%#v' should be less than '%#v' but was not ",
			medium, big)
	}

	if big.Less(small) {
		t.Errorf("'%#v' should be greater than '%#v' but was not ",
			big, small)
	}

	if big.Less(medium) {
		t.Errorf("'%#v' should be greater than '%#v' but was not ",
			big, medium)
	}

	if medium.Less(small) {
		t.Errorf("'%#v' should be greater than '%#v' but was not ",
			medium, small)
	}
}

func TestWeightGraphFromFile(t *testing.T) {
	filename := "data/tinyEWG.txt"
	if testing.Short() {
		t.Skip("Not running test using " + filename)
	}
	fd, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed opening %s, %v", filename, err)
	}

	g, err := ReadWeightGraph(fd)
	if err != nil {
		t.Fatalf("Couldn't read graph from %s, %v", filename, err)
	}

	wantV := 8
	gotV := g.V()

	if gotV != wantV {
		t.Errorf("Vertex count, want %d got %d", wantV, gotV)
	}

	wantE := 16
	gotE := g.E()
	if gotE != wantE {
		t.Errorf("Edge count, want %d got %d", wantE, gotE)
	}

	for _, v := range g.Adj(4) {
		switch v.to {
		case 5:
			continue
		case 7:
			continue
		case 0:
			continue
		case 6:
			continue
		}
		switch v.from {
		case 5:
			continue
		case 7:
			continue
		case 0:
			continue
		case 6:
			continue
		}

		t.Errorf("%d should not be adjacent to %d", v.from, v.to)
	}

	for _, v := range g.Adj(1) {
		switch v.to {
		case 5:
			continue
		case 7:
			continue
		case 2:
			continue
		case 3:
			continue
		}
		switch v.from {
		case 5:
			continue
		case 7:
			continue
		case 2:
			continue
		case 3:
			continue

		}
		t.Errorf("%d should not be adjacent to %d", v.from, v.to)
	}
}
