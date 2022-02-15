package main

func main() {
	treeNode := TreeNode{
		Val: 0,
	}
	deleteNode(&treeNode, 0)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
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

	//如果删除的是最末级节点那么降这个节点值返回nil
	if nodedel.Left == nil && nodedel.Right == nil {
		nodedel = nil
	}
	return root

}

func findNode(root *TreeNode, key int) *TreeNode {
	if root == nil || root.Val == key {
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
