package example_binary_search

import "fmt"

//寻找等于 target 的数
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

//寻找第一个大于/大于等于 target 的数
func findMinUpper(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > target {
			//if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

//寻找第一个小于于/小于等于 target 的数
func findMaxLower(nums []int, target int) int {
	left, right := -1, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] < target {
			//if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	nums1 := []int{5}
	fmt.Println(search(nums1, 5))
	nums := []int{3, 4, 6, 7, 9}
	fmt.Println(findMinUpper(nums, 3))
	fmt.Println(findMinUpper(nums, 4))
	fmt.Println(findMinUpper(nums, 6))
	fmt.Println(findMinUpper(nums, 7))
	fmt.Println(findMinUpper(nums, 9))
	fmt.Println()
	fmt.Println(findMaxLower(nums, 3))
	fmt.Println(findMaxLower(nums, 4))
	fmt.Println(findMaxLower(nums, 6))
	fmt.Println(findMaxLower(nums, 7))
	fmt.Println(findMaxLower(nums, 9))
}
