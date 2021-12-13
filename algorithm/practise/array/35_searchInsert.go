package array

import "fmt"

func searchInsert(nums []int, target int) int {
	//二分查找
	left := 0
	right := len(nums) - 1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3}
	fmt.Println(searchInsert(nums, 4))

}
