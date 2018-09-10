package main

import "fmt"

/*
map

map 映射键到值。

map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var k map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	k = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["benson666"] = Vertex{
		40, -74,
	}
	k["Benson Maxs"] = Vertex{
		50.6666, -111,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(k["Benson Maxs"])
	fmt.Println(m["benson666"])
}
