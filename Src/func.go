package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupport op" + op)
	}

}

func main() {
	fmt.Println(eval(1, 2, "+"))
}
