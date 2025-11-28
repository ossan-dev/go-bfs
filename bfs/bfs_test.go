package bfs_test

import (
	"testing"

	"github.com/ossan-dev/go-bfs/bfs"
	"github.com/stretchr/testify/require"
)

var g *bfs.Graph = &bfs.Graph{}

func BenchmarkWalkFromNode(b *testing.B) {
	b.ReportAllocs()
	expected := []string{"operating systems", "data structures", "computer organization", "discrete math", "intro to programming"}
	for b.Loop() {
		b.StopTimer()
		courses := make([]string, numberOfCourses)
		b.StartTimer()
		// g.WalkFromNodeSlice("networks", &courses) // slices
		g.WalkFromNodeMap("networks", &courses) // maps
		require.ElementsMatch(b, expected, courses)
	}
}
