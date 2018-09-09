package main

import "fmt"

/*
定义函数
函数名在前，返回值类型在后
参数名在前，返回值类型在后
*/
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
