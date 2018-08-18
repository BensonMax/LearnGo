package main

import "fmt"

/*
Slice 本身没有数据，是对底层array的一个view
*/
func undateslice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[:]
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] =", arr[2:])
	s2 := arr[2:6]
	fmt.Println("arr[:] =", arr[:])

	fmt.Println("After update")
	undateslice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println(s2)
	fmt.Println(s)

}
