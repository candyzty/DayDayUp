package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode
	fmt.Println(root)
	root.value = 23
	root.left = &treeNode{}
	root.right = &treeNode{3, nil, nil}
	root.right.left = new(treeNode)

}
