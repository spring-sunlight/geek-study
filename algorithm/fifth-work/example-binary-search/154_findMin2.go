package example_binary_search

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		//fmt.Println("index:", left, right, mid)
		//fmt.Println("nums:", nums[left:right+1])
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	//fmt.Println("=================")
	return nums[right]
}
