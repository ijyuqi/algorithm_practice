package algorithm

import "fmt"

// 二叉树节点

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

// 二叉树

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(data string) {
	if t.Root == nil {
		t.Root = &Node{Data: data, Left: nil, Right: nil}
	} else {
		t.InsertNode(t.Root, data)
	}
}

func (t *BinaryTree) InsertNode(node *Node, data string) {
	if node.Left == nil {
		node.Left = &Node{Data: data, Left: nil, Right: nil}
	} else if node.Right == nil {
		node.Right = &Node{Data: data, Left: nil, Right: nil}
	} else {
		t.InsertNode(node.Left, data)
	}
}

func (t *BinaryTree) PreorderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%s ", node.Data)
		t.PreorderTraversal(node.Left)
		t.PreorderTraversal(node.Right)
	}
}

// 中序遍历

func (t *BinaryTree) InorderTraversal(node *Node) {
	if node != nil {
		t.InorderTraversal(node.Left)
		fmt.Printf("%s ", node.Data)
		t.InorderTraversal(node.Right)
	}
}

// 中序遍历

func (t *BinaryTree) PostorderTraversal(node *Node) {
	if node != nil {
		t.PostorderTraversal(node.Left)
		t.PostorderTraversal(node.Right)
		fmt.Printf("%s ", node.Data)
	}
}
