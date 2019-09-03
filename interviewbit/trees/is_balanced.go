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

func initTree(inputSlice []int) *TreeNode {
	root := &TreeNode{
		data: inputSlice[0],
	}
	lastParent := root
	for _, i := range inputSlice[1:] {
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	return max(height(root.left), height(root.right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	imbalance := height(root.left) - height(root.right)
	fmt.Printf("Imbalance at node %d: %d\n", root.data, imbalance)
	if imbalance > 1 || imbalance < -1 {
		return false
	}
	return isBalanced(root.left) && isBalanced(root.right)
}

func main() {
	t := initTree()
	inOrder(t)
	fmt.Println(isBalanced(t))
}
