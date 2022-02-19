package main

import "container/list"

func maxDepth(root *TreeNode) int {

	return bfs104(root)
}

func bfs104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
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
