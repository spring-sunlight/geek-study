package array

func getRow(rowIndex int) []int {
	ans := []int{}

	ans = append(ans, 1)

	for i := 1; i <= rowIndex; i++ {
		//fmt.Println("ans:", ans)
		temp := make([]int, i+1)
		temp[0] = 1
		temp[len(temp)-1] = 1
		for j := 1; j < i; j++ {
			temp[j] = ans[j] + ans[j-1]
		}
		//fmt.Println("temp:", temp)
		ans = temp
	}
	return ans
}
