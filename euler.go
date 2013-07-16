package graph

func HasEulerTour(g Ungraph) bool {

	isOdd := func(n int) bool { return n%2 != 0 }

	for v := 0; v < g.V(); v++ {
		if isOdd(Degree(g, v)) {
			return false
		}
	}
	return true
}

func EulerTour(g Ungraph) []int {
	panic("not implemented")
	return []int{}
}

func HasEulerCircuit(g Ungraph) bool {
	panic("not implemented")
	return false
}

func EulerCircuit(g Ungraph) []int {
	panic("not implemented")
	return []int{}
}
