package main

import "fmt"

func main() {
	var array1 [5]int
	array2 := [3]int{1, 3, 5}
	array3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool
	fmt.Println(array1, array2, array3, grid)

	//遍历数组1 推荐
	for i := range array3 {
		fmt.Println(i)
	}

	//遍历数组2
	for i := 0; i < len(array3); i++ {
		fmt.Println(i)
	}
}
