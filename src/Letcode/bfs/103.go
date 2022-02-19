package main

import (
	"container/list"
)

func bfs103(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
	result := make([][]int, 0, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		curlevel := make([]int, 0, 0)
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			curlevel = append(curlevel, node.Val)
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		if level%2 == 0 {
			sortLevel := make([]int, 0, 0)
			for i := len(curlevel) - 1; i >= 0; i-- {
				sortLevel = append(sortLevel, curlevel[i])
			}
			result = append(result, sortLevel)
		} else {
			result = append(result, curlevel)
		}

		level++
	}
	return result
}
