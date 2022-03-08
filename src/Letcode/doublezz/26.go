/**
* @Author: Gosin
* @Date: 2022/3/1 22:47
 */

package doublezz

func removeDuplicates(nums []int) int {

	if nums == nil {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(nums)-1; fast++ {
		if nums[slow] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
