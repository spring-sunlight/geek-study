package example_binary_search

//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//寻找第一个大于等于 target 的数
//寻找最后一个小于等于 target 的数
func searchRange(nums []int, target int) []int {
	ans := make([]int, 2)
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	ans[0] = right
	left, right = -1, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	ans[1] = right
	if ans[0] > ans[1] {
		return []int{-1, -1}
	}
	return ans
}
