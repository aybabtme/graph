package unionfind

import (
	"testing"
)

var (
	tinyUF = []struct{ from, to int }{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
	}

	expectedCount = 2
)

func TestUnionFindMatchesKnownModel(t *testing.T) {
	uf := BuildUF(10)
	for _, pair := range tinyUF {
		uf.Union(pair.from, pair.to)

		if !uf.Connected(pair.from, pair.to) ||
			!uf.Connected(pair.to, pair.from) {

			t.Errorf("Union (%d,%d) was not recorded", pair.from, pair.to)
		}
	}

	actualCount := uf.Count()

	if expectedCount != actualCount {
		t.Errorf("Should have counted %d components, but was %d",
			expectedCount, actualCount)
	}

	count := make(map[int]bool)
	for i := 0; i < 10; i++ {
		count[uf.Find(i)] = true
	}

	if len(count) != expectedCount {
		t.Errorf("Expected %d types of ids, but got %d and ids=%v",
			expectedCount, len(count), count)
	}

	for i := 0; i < 10; i++ {
		uf.Union(i, i)
		if !uf.Connected(i, i) {
			t.Errorf("Not connected to itself: %d", i)
		}
	}
}
