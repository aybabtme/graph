package mst

import (
	"github.com/aybabtme/graph"
	"os"
	"testing"
)

type mstBuilder func(*graph.WeightGraph) MST

type lazyWeightGraph func() *graph.WeightGraph

var (
	tiny   lazyWeightGraph
	medium lazyWeightGraph
	large  lazyWeightGraph
)

func init() {

	var tinyG *graph.WeightGraph
	tiny = func() *graph.WeightGraph {
		if tinyG == nil {
			tinyG = readGraph("../data/tinyEWG.txt")
		}
		return tinyG
	}

	var mediumG *graph.WeightGraph
	medium = func() *graph.WeightGraph {
		if mediumG == nil {
			mediumG = readGraph("../data/mediumEWG.txt")
		}
		return mediumG
	}

	var largeG *graph.WeightGraph
	large = func() *graph.WeightGraph {
		if largeG == nil {
			largeG = readGraph("../data/largeEWG.txt")
		}
		return largeG
	}

}

func readGraph(filename string) *graph.WeightGraph {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := graph.ReadWeightGraph(f)
	if err != nil {
		panic(err)
	}
	return &g
}

func Benchmark_Kruskal_TinyEWG(b *testing.B)   { benchmarkGraph(BuildKruskalMST, tiny, b) }
func Benchmark_Kruskal_MediumEWG(b *testing.B) { benchmarkGraph(BuildKruskalMST, medium, b) }
func Benchmark_Kruskal_LargeEWG(b *testing.B)  { benchmarkGraph(BuildKruskalMST, large, b) }

func Benchmark_Prime_TinyEWG(b *testing.B)   { benchmarkGraph(BuildLazyPrimMST, tiny, b) }
func Benchmark_Prime_MediumEWG(b *testing.B) { benchmarkGraph(BuildLazyPrimMST, medium, b) }
func Benchmark_Prime_LargeEWG(b *testing.B)  { benchmarkGraph(BuildLazyPrimMST, large, b) }

func benchmarkGraph(mstB mstBuilder, graphLoader lazyWeightGraph, b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	_ = mstB(graphLoader())
	b.StopTimer()
}
