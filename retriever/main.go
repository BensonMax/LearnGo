package main

import (
	"LearnGo/retriever/mock"
	real2 "LearnGo/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake baidu.com"}
	r = real2.Retriever{}
	fmt.Println(download(r))
}
