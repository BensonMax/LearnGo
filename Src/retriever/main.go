package main

import "fmt"

type a struct {
	x int
}
type b struct {
	x int
}

func (*a) Get(url string) string {

	return "sdf "
}
func (*b) Get(url string) string {

	return "sdf "
}

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.imooc.com")
}

func main() {
	//var r Retiever
	var xy = a{1}
	var yx = b{2}
	//fmt.Println(download(r))
	fmt.Println(download(&xy))
	fmt.Println(download(&yx))
}
