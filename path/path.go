package path

// PathFinder holds information regarding the paths in a graph from a source
type PathFinder interface {
	// HasPathTo tells whether there is a path between the source and the
	// destination
	HasPathTo(destination int) bool
	// PathTo returns the path from the source to the destination
	PathTo(destination int) []int
}
