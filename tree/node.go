package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	//返回局部变量地址给别人用，不需要知道堆还是栈，系统会自动回收
	return &treeNode{value: value}
}
//在函数名的前面增加了一个接收者
//实际上如果将接收者放在形参也差不多能实现差不多的效果
func (node treeNode) print() {
	fmt.Println(node.value, " ")
}
//指针模拟引用传递，传递地址。但是node.value的调用方式没有改变
func (node *treeNode) setValue(value int){
	node.value = value
}


func (node *treeNode) traverse() {
	//只要判断了就行
	if node == nil {
		return
	}
	//这里java之类的是一定要判断是否为空的，不然没法用left的方法，但是golang就可以，只需要内部对nil值做处理了就行
	node.left.traverse()
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//new(type)返回的是地址
	root.right.left= new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(root)
	fmt.Println(nodes)
	root.print()
	root.right.left.setValue(4)
	root.right.left.print()

	//print是值接收者，setvalue是指针接收者，但是都可以直接.调用
	root.print()
	//这里虽然是值调用，但是实际上接收者是指针类型，因此可以改变原来的数据。
	//就是说能不能改变是由接受者类型来决定的，而不是实参。
	root.setValue(100)
	root.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	root.traverse()

}
