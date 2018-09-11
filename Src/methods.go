package main

/*
Go 没有类。然而，仍然可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//方法接收者
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Add() float64 {
	return math.Sqrt(v.X * v.X)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Add())
}
