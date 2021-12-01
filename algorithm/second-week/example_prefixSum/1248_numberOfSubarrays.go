package example_prefixSum

func numberOfSubarrays(nums []int, k int) int {
	//1.前缀和数组长度比原数组多一个,即首位多一个 0
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = nums[i-1]%2 + prefixSum[i-1]
	}
	//2.转化成两数之差等于k的问题,数组:prefixSum,target:k
	numsMap := make(map[int]int)
	ans := 0
	for i := 0; i < len(prefixSum); i++ {
		if _, ok := numsMap[prefixSum[i]-k]; ok {
			ans += numsMap[prefixSum[i]-k]
		}
		numsMap[prefixSum[i]]++
	}
	return ans
}
