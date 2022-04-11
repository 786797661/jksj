package juz

import "container/list"

func networkDelayTime(times [][]int, n int, k int) int {
	queue := list.New()
	//记录初始节点
	for _, time := range times {
		if k == time[0] {
			queue.PushBack(time)
		}
	}
	if queue.Len() < 0 {
		return -1
	}
	for queue.Len() > 0 {
		node := queue.Front().Value.([]int)
		for _, time := range times {
			if node[1] == time[0] {
				queue.PushBack(time)
			}
		}

	}
}
