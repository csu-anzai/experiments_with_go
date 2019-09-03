package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// gets first node that has a nil child
func getNextParent(root *TreeNode) (*TreeNode, string) {
	if root.left == nil {
		return root, "left"
	} else if root.right == nil {
		return root, "right"
	} else {
		if leftTree, side := getNextParent(root.left); leftTree != nil {
			return leftTree, side
		}
		if rightTree, side := getNextParent(root.right); rightTree != nil {
			return rightTree, side
		}
	}
	return nil, ""
}

func initTree() *TreeNode {
	root := &TreeNode{
		data: 1,
	}
	lastParent := root
	for _, i := range []int{2, 3, 4, 5, 6, 7} {
		nextParent, side := getNextParent(lastParent)
		if side == "left" {
			nextParent.left = &TreeNode{
				data: i,
			}
		} else {
			nextParent.right = &TreeNode{
				data: i,
			}
		}
	}
	return root
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	} else {

		//if root.left != nil && root.right != nil {
		//	fmt.Println("2 children")
		//} else if root.left != nil {
		//	fmt.Println("left child only")
		//} else if root.right != nil {
		//	fmt.Println("right child only")
		//}
		inOrder(root.left)
		fmt.Println(root.data)
		inOrder(root.right)
	}

}

func main() {
	t := initTree()
	inOrder(t)
}
