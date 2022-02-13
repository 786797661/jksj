package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	//[5,4,6,null,null,3,7]

	node := TreeNode{Val: 5}
	node1 := TreeNode{Val: 4}
	node2 := TreeNode{Val: 6}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 7}
	node.Left = &node1
	node.Right = &node2
	node2.Left = &node3
	node2.Right = &node4
	fmt.Println(isValidBST(&node))
}
func isValidBST(root *TreeNode) bool {
	return helper(root)
}

//func isValidBST(root *TreeNode) bool {
//	return helper(root, math.MinInt64, math.MaxInt64)
//}

//func helper(root *TreeNode, lower, upper int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val <= lower || root.Val >= upper {
//		return false
//	}
//	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
//}

func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flagLeft := false
	flagRight := false
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Left.Val < root.Val {
		flagLeft = true
	}
	if root.Right != nil && root.Right.Val > root.Val {
		flagRight = true
	}
	if flagLeft && flagRight {
		return helper(root.Left) && helper(root.Right)
	}
	return false
}
