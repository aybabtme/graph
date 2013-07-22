package path

import (
	"github.com/aybabtme/graph"
	"testing"
)

var (
	// We use this crazy digraph, that's also a DAG, for our TC testing
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
	tcGraphEdges = []struct{ from, to int }{
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
	// We know the former are reachable from the latter
	tcReachableExpected = []map[int]bool{
		/*  0 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true},
		/*  1 */ {1: true},
		/*  2 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true},
		/*  3 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true},
		/*  4 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true},
		/*  5 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true},
		/*  6 */ {
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true,
			// Can't reach 7
			8: true, 9: true, 10: true, 11: true, 12: true,
		},
		/*  7 */ {
			// Can reach all the others
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true,
		},
		/*  8 */ {0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true,
			// Can't reach 7
			8: true, 9: true, 10: true, 11: true, 12: true,
		},
		/*  9 */ {
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true,
			// Can't reach 6, 7 and 8
			9: true, 10: true, 11: true, 12: true,
		},
		/* 10 */ {
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true,
			// Can't reach 6, 7 and 8
			9: true, 10: true, 11: true, 12: true,
		},
		/* 11 */ {
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true,
			// Can't reach 6, 7 and 8
			9: true, 10: true, 11: true, 12: true,
		},
		/* 12 */ {
			0: true, 1: true, 2: true, 3: true, 4: true, 5: true,
			// Can't reach 6, 7 and 8
			9: true, 10: true, 11: true, 12: true,
		},
	}
)

func TestTransitiveClosureMatchesOutput(t *testing.T) {
	di := graph.NewDigraph(13)
	for _, edge := range tcGraphEdges {
		di.AddEdge(edge.from, edge.to)
	}

	tc := BuildTransitiveClosure(di)

	for v := 0; v < di.V(); v++ {
		for w := 0; w < di.V(); w++ {
			expected := tcReachableExpected[v][w]
			actual := tc.Reachable(v, w)
			if expected != actual {
				t.Errorf("Expected reachability from %d to %d to be %v but was %v",
					v, w, expected, actual)
			}
		}
	}
}
