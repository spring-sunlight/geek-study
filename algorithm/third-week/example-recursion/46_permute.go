package example_recursion

import "fmt"

func permute(nums []int) [][]int {
	ans := [][]int{}
	share := []int{}
	used := make([]int, len(nums))

	var recur func(i int)
	recur = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int{}, share...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if used[j] == 0 {
				share = append(share, nums[j])
				used[j] = 1
				recur(i + 1)
				used[j] = 0
				share = share[:len(share)-1]
			}
		}
	}
	recur(0)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
