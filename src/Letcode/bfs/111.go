package main

import "container/list"

func minDepth(root *TreeNode) int {
	return bfs111(root)
}
func bfs111(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)

	level := 1

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return level
			}
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}
