package main

import (
	"fmt"
	"rpc/algorithm"
)

func main() {

	tree := algorithm.BinaryTree{}
	tree.Insert("A")
	tree.Insert("B")
	tree.Insert("C")
	tree.Insert("D")
	tree.Insert("E")
	tree.Insert("F")

	fmt.Print("Inorder traversal: ")
	tree.InorderTraversal(tree.Root)
	fmt.Println()

	arr := []int{1, 2, 5, 3, 0, 8, 7, 4, 6}
	fmt.Println("排序前", arr)

	algorithm.BubbleSort(arr)
	fmt.Println("排序后", arr)
}
