package goplayround

import (
	"testing"
)

var s = createDenseMatrixSpawner(1024 * 8)
var r float64
func benchmarkTwoNorm(scale int, b *testing.B) {
	m := s(scale)
	b.ResetTimer()
	for i := 0; i < b.N ; i++ {
		r = m.TwoNorm()
	}
}

func benchmarkTwoNormParallel(scale ,workers int, b *testing.B ) {
	m := s(scale)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = TwoNormParallel(m, workers)
	}
}

func BenchmarkTwoNorm128(b *testing.B) { benchmarkTwoNorm(128, b)}
func BenchmarkTwoNormParallel128_4(b *testing.B) { benchmarkTwoNormParallel(128, 4, b)}
func BenchmarkTwoNormParallel128_8(b *testing.B) { benchmarkTwoNormParallel(128, 8, b)}
func BenchmarkTwoNorm1024(b *testing.B) { benchmarkTwoNorm(1024, b)}
func BenchmarkTwoNormParallel1024_4(b *testing.B) { benchmarkTwoNormParallel(1024, 4, b)}
func BenchmarkTwoNormParallel1024_8(b *testing.B) { benchmarkTwoNormParallel(1024, 8, b)}
func BenchmarkTwoNorm10248(b *testing.B) { benchmarkTwoNorm(1024 * 8, b)}
func BenchmarkTwoNormParallel10248_4(b *testing.B) { benchmarkTwoNormParallel(1024 * 8, 4, b)}
func BenchmarkTwoNormParallel10248_8(b *testing.B) { benchmarkTwoNormParallel(1024 * 8, 8, b)}
