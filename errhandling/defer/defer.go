package main

import (
	"LearnGo/functional/fib"
	"bufio"
	"fmt"
	"os"
)

/*
defer 1.当函数退出时执行 2.先进后出
*/

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

// 写入文件
func writeFile(filename string) {
	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	//关闭文件
	defer file.Close()
	//
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
