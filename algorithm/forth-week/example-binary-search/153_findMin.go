package example_binary_search

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		//fmt.Println("left:", left, ",right:", right, ",mid:", mid)
		if nums[mid] <= nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
