package example_array

import "fmt"

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)

}

func testMoveZeroes() {
	nums := []int{1, 0, 2, 3, 12, 0}
	moveZeroes(nums)
}
