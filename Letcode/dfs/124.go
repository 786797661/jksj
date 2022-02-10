package main

import "math"

func main() {

}

func maxPathSum(root *TreeNode) int {
	//return maxThreeNode(root)
	var maxGin func(*TreeNode) int
	sumMax := math.MaxInt32
	/**
	获取最大贡献值
	*/
	maxGin = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//当前节点的最大贡献值=当前节点值+max(左，右)
		//当前节点的路径和=当前节点值+左最大贡献值+右最大贡献值

		leftGx := max(maxGin(node.Left), 0)
		rightGx := max(maxGin(node.Right), 0)

		sumTemp := leftGx + rightGx + node.Val

		sumMax = max(sumMax, sumTemp)
		return node.Val + max(leftGx, rightGx)
		//return  max(maxThreeNode(root.Left),maxThreeNode(root.Right)) + root.Val
	}

	maxGin(root)
	return sumMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
