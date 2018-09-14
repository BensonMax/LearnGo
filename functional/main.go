package main

import (
	"LearnGo/functional/fib"
	"fmt"
)

func main() {
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("index%d %d\n",
			i, f())
	}
}
