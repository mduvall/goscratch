package structures

import (
	"fmt"
)

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func CreateTree(vals []int) *Tree {
	tree := Tree{}

	for _, num := range vals {
		node := TreeNode{Left: nil, Right: nil, Val: num}
		tree.Insert(&node)
	}

	return &tree
}

func (t *Tree) Insert(n *TreeNode) {
	if t.Root == nil {
		t.Root = n
		return
	}

	root := t.Root
	for {
		if n.Val <= root.Val {
			if root.Left != nil {
				root = root.Left
			} else {
				root.Left = n
				break
			}
		}

		if n.Val > root.Val {
			if root.Right != nil {
				root = root.Right
			} else {
				root.Right = n
				break
			}
		}
	}
}

// Simple tree traversal
func Print(n *TreeNode) {
	if n == nil {
		return
	}

	Print(n.Left)
	Print(n.Right)
	fmt.Println(n.Val)
}

func (t *Tree) Search(i int) *TreeNode {
	root := t.Root

	for root != nil {
		if i == root.Val {
			return root
		}

		if i > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return nil
}

func (t *Tree) Min() int {
	root := t.Root

	for root.Left != nil {
		root = root.Left
	}

	return root.Val
}

func (t *Tree) Max() int {
	root := t.Root

	for root.Right != nil {
		root = root.Right
	}

	return root.Val
}
