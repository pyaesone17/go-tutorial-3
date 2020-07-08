package benchmarks

import "fmt"

// Fib ...
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop")
	}
}

func LoopWhile() {
	i := 10
	for i < 10 {
		fmt.Println("Loop")
		i++
	}
}
