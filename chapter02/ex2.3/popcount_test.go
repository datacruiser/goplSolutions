// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
/*
⇒  go test -bench=.
ggoos: darwin
goarch: amd64
pkg: github.com/jijiwhywhy/goplSolutions/chapter02/ex2.3
BenchmarkPopCount-8             1000000000               0.291 ns/op
BenchmarkLoopPopCount-8         72501686                16.5 ns/op
PASS
ok      github.com/jijiwhywhy/goplSolutions/chapter02/ex2.3     2.845s
*/
package popcount_test

import (
	"testing"

	"github.com/jijiwhywhy/goplSolutions/chapter02/ex2.3"
)

// -- Alternative implementations --

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.LoopPopCount(0x1234567890ABCDEF)
	}
}
