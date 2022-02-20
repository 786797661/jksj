/**
* @Author: Gosin
* @Date: 2022/2/19 23:06
 */

package main

import (
	"container/list"
)

func largestValues(root *TreeNode) []int {
	return bfs515(root)
}

func bfs515(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	levelMaxs := make([]int, 0, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		levelMax := queue.Front().Value.(*TreeNode).Val
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			if levelMax <= node.Val {
				levelMax = node.Val
			}
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
		levelMaxs = append(levelMaxs, levelMax)
	}
	return levelMaxs
}
