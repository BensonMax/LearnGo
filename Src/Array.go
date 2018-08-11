package main

import "fmt"

/*
数组是值类型
数组是值类型传递，调用函数会拷贝数组，与其他语言不一样
go语言一般不直接数组，也不使用数组的指针；使用切片
*/
func printarray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

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

	//求数组value 求和 , 通过_略过变量
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)

	printarray(array3)
	printarray(array1)
	//printarray(array2)  数组是值类型 arr 定义是arr[5] 而arr2 是 arr[3]
}
