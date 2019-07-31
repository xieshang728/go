package tree

import "fmt"

type treeNode struct{
	value int
	left,right *treeNode
}


func CreateTreeNode(val int) *treeNode{
	return &treeNode{value :val}
}

func (node *treeNode) SetValue(value int){
	node.value = value
}

func (node *treeNode) print(){
	fmt.Println(node)
}


func main(){
	node := treeNode{value:5}
	fmt.Println(node)
	node.SetValue(120)
	fmt.Println(node)
	fmt.Println()
	node.print()
}
