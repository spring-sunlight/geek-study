package practise

import "sort"

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	share := []int{}
	used := make([]int, len(nums))
	sort.Ints(nums)

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int{}, share...))
		}

		for j := 0; j < len(used); j++ {
			if j >= 1 && used[j-1] == 1 && nums[j] == nums[j-1] {
				continue
			}
			if used[j] == 0 {
				share = append(share, nums[j])
				used[j] = 1
				dfs(i + 1)
				used[j] = 0
				share = share[:len(share)-1]
			}
		}
	}
	dfs(0)
	return ans
}
