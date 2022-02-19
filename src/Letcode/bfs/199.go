package main

import "container/list"

func rightSideView(root *TreeNode) []int {
	return bfs199(root)
}

func bfs199(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	result := make([]int, 0, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if i == size-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}
