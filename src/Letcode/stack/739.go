/**
* @Author: Gosin
* @Date: 2022/2/19 23:29
 */

package stack

import (
	"container/list"
)

func dailyTemperatures(temperatures []int) []int {
	return stack739(temperatures)
}

func stack739(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return nil
	}
	stack := list.New()

	res := make([]int, len(temperatures), cap(temperatures))
	for i, temperature := range temperatures {
		for stack.Len() > 0 {
			index := stack.Front().Value.(int)
			if temperatures[index] < temperature {
				stack.Remove(stack.Front())
				res[index] = i - index
			} else {
				break
			}
		}
		stack.PushFront(i)
	}
	//for stack.Len()>0 {
	//	index:=stack.Front().Value.(int)
	//	stack.Remove(stack.Front())
	//	res[index]=0
	//}
	return res
}
