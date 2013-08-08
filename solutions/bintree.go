package main

/*
Write a binary search tree with insert, find,
and three different traversal print functions
*/
import (
	"fmt"
)

func main() {
	t := new(TreeNode)
	t.val = 45
	t.Insert(3)
	t.Insert(123)
	t.Insert(6)
	t.Insert(36)
	t.Insert(48)
	t.Insert(99)
	t.Insert(1)
	t.Insert(14)
	t.Insert(77)

	fmt.Println("In Order:")
	t.PrintInOrder()
	fmt.Println("\n\nPre Order:")
	t.PrintPreOrder()
	fmt.Println("\n\nPost Order:")
	t.PrintPostOrder()

}

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	val int
}

func (t *TreeNode) Insert(n int) {
	if n > t.val {
		if t.right == nil {
			t.right = &TreeNode{nil,nil,n}
		} else {
			t.right.Insert(n)
		}
	} else {
		if t.left == nil {
			t.left = &TreeNode{nil,nil,n}
		} else {
			t.left.Insert(n)
		}
	}
}

func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}
	t.left.PrintInOrder()
	fmt.Println(t.val)
	t.right.PrintInOrder()
}

func (t *TreeNode) PrintPreOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.val)
	t.left.PrintPreOrder()
	t.right.PrintPreOrder()
}

func (t *TreeNode) PrintPostOrder() {
	if t == nil {
		return
	}
	t.left.PrintPostOrder()
	t.right.PrintPostOrder()
	fmt.Println(t.val)
}

