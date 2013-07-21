package path

// PathFinder holds information regarding the paths in a graph from a source
type PathFinder interface {
	// HasPathTo tells whether there is a path between the source and the
	// destination
	HasPathTo(destination int) bool
	// PathTo returns the path from the source to the destination
	PathTo(destination int) []int
}

func reverse(s []int) []int {
	var opposite int
	for i := 0; i < len(s)/2; i++ {
		opposite = len(s) - 1 - i
		s[i], s[opposite] = s[opposite], s[i]
	}
	return s
}
