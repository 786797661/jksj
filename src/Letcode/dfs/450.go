package main

//func main() {
//	treeNode := TreeNode{
//		Val: 0,
//	}
//	deleteNode(&treeNode, 0)
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNodeTest(root *TreeNode, key int) *TreeNode {
	nodedel := findNode(root, key)
	if nodedel == nil {
		return root
	}
	if nodedel.Right != nil {
		nodedel.Val = nodedel.Right.Val
		if nodedel.Right.Right == nil {
			nodedel.Right = nodedel.Right.Right
		} else {
			nodedel.Right = nil
		}
	} else if nodedel.Left != nil {
		nodedel.Val = nodedel.Left.Val
		if nodedel.Left.Left == nil {
			nodedel.Left = nodedel.Left.Left
		} else {
			nodedel.Left = nil
		}
	}
	if nodedel == root {
		return nil
	} else {
		nodedel = nil
	}
	return root
}

func findNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		//if root.Left == nil && root.Right == nil {
		//
		//}
		return root
	}

	left := findNode(root.Left, key)
	right := findNode(root.Right, key)
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}
