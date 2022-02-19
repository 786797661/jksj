package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(level int, root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := list.New()
	//存入根节点
	queue.PushBack(root)
	result := make([][]int, 0, 0)

	//遍历取出
	for queue.Len() > 0 {
		levelsize := queue.Len()
		cuerlevel := make([]int, 0, 0)

		for i := 0; i < levelsize; i++ {
			node := queue.Front().Value.(TreeNode)
			cuerlevel = append(cuerlevel, node.Val)
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, cuerlevel)
	}
	return result
}
