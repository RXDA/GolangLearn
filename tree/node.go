package main

import "fmt"

type treeNode struct{
	value int
	left,right *treeNode
}
//类似于构造函数
func (node treeNode) print(){//括号里是接收者,相当于this
	fmt.Print(node.value,"\n")
}

func (node treeNode) setValue(value int){
	node.value=value
}

func (node *treeNode) setValueTrue(value int){
	node.value=value
}
//不论是指针接收者还是值接收者都可直接调用

func createNode(value int) *treeNode{
	return &treeNode{value:value} //可以返回一个局部变量的地址
}

//type mineTreeNode struct {
//	value int
//	left,right mineTreeNode
//}





func main() {
	var root treeNode
	//fmt.Println(root)

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right=&treeNode{5,nil,nil}
	root.right.left= new(treeNode)
	root.left.right=createNode(2)
	//fmt.Println(root)

	root.right.left.setValue(1)//没有改变值
	root.right.left.print()

	root.right.left.setValueTrue(6)//改变了值
	root.right.left.print()
	//var root1 mineTreeNode
	//fmt.Println(root1)
	//
	//root1 = mineTreeNode{value:3}
	//root1.left=mineTreeNode{}
	//root1.right=mineTreeNode{5,nil,nil}
	//fmt.Println(root1)

	nodes:=[]treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
	root.print()
}
