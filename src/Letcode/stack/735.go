/**
* @Author: Gosin
* @Date: 2022/2/20 22:37
 */

package stack

import (
	"container/list"
)

func asteroidCollision(asteroids []int) []int {
	if asteroids == nil {
		return nil
	}
	return stack735(asteroids)
}

func stack735(asteroids []int) []int {
	stack := list.New()
	for i, asteroid := range asteroids {
		isDes := false
		for stack.Len() > 0 {
			index := stack.Front().Value.(int)
			//栈顶向左就不会发生碰撞
			if asteroids[index] < 0 || asteroids[index]*asteroid > 0 {
				break
			}
			if asteroids[index]*asteroid < 0 {
				//不同方向的
				//如果站里是负数向左就不会碰撞否则比较两者大小,由于前两个
				if asteroid*-1 > asteroids[index] {
					stack.Remove(stack.Front())
				} else if asteroid*-1 == asteroids[index] {
					//相同质量的都会被消亡
					stack.Remove(stack.Front())
					isDes = true
					break
				} else {
					isDes = true
					break
				}

			}
		}
		if !isDes {
			stack.PushFront(i)
		}
	}
	res := []int{}
	for stack.Len() > 0 {
		res = append(res, asteroids[stack.Back().Value.(int)])
		stack.Remove(stack.Back())
	}
	return res
}
