package main

import (
	"testing"

	"github.com/luismiguel010/match/match"
)

func BenchmarkMatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		match.TotalMatch()
	}
}
