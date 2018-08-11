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
	//进阶 返回key value
	for i, v := range array3 {
		fmt.Println(i, v)
	}
	//进阶2 只返回value
	for _, v := range array3 {
		fmt.Println(v)
	}

	//遍历数组2
	for i := 0; i < len(array3); i++ {
		fmt.Println(i)
	}

	//求数组中最大的 key 和 value
	numbers := [6]int{1, 3, 5, 7, 9, 11}

	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
		fmt.Println(maxi, maxValue)
	}
}
