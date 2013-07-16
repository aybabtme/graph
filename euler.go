package graph

func HasEulerTour(g Graph) bool {

	isOdd := func(n int) bool { return n%2 != 0 }

	for v := 0; v < g.V(); v++ {
		if isOdd(Degree(g, v)) {
			return false
		}
	}
	return true
}

func EulerTour(g Graph) []int {
	panic("not implemented")
	return []int{}
}

func HasEulerCircuit(g Graph) bool {
	panic("not implemented")
	return false
}

func EulerCircuit(g Graph) []int {
	panic("not implemented")
	return []int{}
}
