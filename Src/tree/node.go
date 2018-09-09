package tree

import "fmt"

/*
面向对象
1.go语言仅支持封装，不支持继承和多态
2.go语言没有class，只有struct
*/

//创建结构体，type指定类型
type Node struct {
	Value       int
	Left, Right *Node
}

//定义print方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" + "node.Ignored")
		return
	}
	node.Value = value
}

//定义Node遍历函数
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}
