package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//示例 2:
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nums := []int{5}
	search(nums, 5)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 0 {
		if nums[0] == target {
			return 0
		}
	}
	low, high := 0, len(nums)-1

	for low < high {
		var mid = (high-low)/2 + low
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
