/**
* @Author: Gosin
* @Date: 2022/2/19 23:20
 */

package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	return bfs429(root)
}
func bfs429(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	res := make([][]int, 0, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		levelNode := make([]int, 0, 0)
		for i := 0; i < size; i++ {
			dnode := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())
			if dnode.Children != nil {
				for _, child := range dnode.Children {
					if child != nil {
						queue.PushBack(child)
					}
				}
			}
			levelNode = append(levelNode, dnode.Val)
		}
		res = append(res, levelNode)
	}
	return res
}
