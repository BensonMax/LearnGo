package main

import (
	"LearnGo/retriever/mock"
	real2 "LearnGo/retriever/real"
	"fmt"
	"time"
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
	fmt.Printf("%T %v\n", r, r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))

	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("content", v.Content)
	case *real2.Retriever:
		fmt.Println("Useragent", v.UserAgent)
	}
}
