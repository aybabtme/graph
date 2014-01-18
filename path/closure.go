package path

import (
	"github.com/aybabtme/graph"
)

// TransitiveClosure answers queries of the form "Is there a directed path
// from vertex v to vertex w?".
type TransitiveClosure struct {
	paths []PathFinder
}

// BuildTransitiveClosure builds a model of all-pair reachable vertices using
// depth-first search path finders.
func BuildTransitiveClosure(di *graph.Digraph) *TransitiveClosure {
	tc := &TransitiveClosure{paths: make([]PathFinder, di.V())}

	for v := 0; v < di.V(); v++ {
		tc.paths[v] = BuildDFS(di, v)

	}
	return tc
}

// Reachable tells if vertex w is reachable from vertex v.
func (t *TransitiveClosure) Reachable(v, w int) bool {
	return t.paths[v].HasPathTo(w)
}
