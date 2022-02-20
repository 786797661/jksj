/**
* @Author: Gosin
* @Date: 2022/2/20 23:18
 */

package stack

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	return stack496(nums1, nums2)
}

func stack496(nums1 []int, nums2 []int) []int {
	stack := list.New()
	res := []int{}
	for _, value := range nums1 {
		for i := len(nums2) - 1; i >= 0; i-- {
			//凡是大于本身的放入站中
			if value < nums2[i] {
				if stack.Len() > 0 {
					stack.Remove(stack.Front())
				}
				stack.PushFront(i)
			}
			//找到本身，如果站有数据就出栈，没有数据就为-1
			if value == nums2[i] {
				if stack.Len() > 0 {
					res = append(res, nums2[stack.Front().Value.(int)])
					stack.Remove(stack.Front())
				} else {
					res = append(res, -1)
				}
				break
			}
		}
	}
	return res
}
