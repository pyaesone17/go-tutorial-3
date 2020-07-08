package benchmarks

import "testing"

// // from fib_test.go
// func BenchmarkFib10(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		Fib(10)
// 	}
// }

func BenchmarkLoop10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Loop()
	}
}

func BenchmarkLoopWhile10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LoopWhile()
	}
}
