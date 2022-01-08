package array

import "fmt"

func generate(numRows int) [][]int {
	ans := [][]int{}

	ans = append(ans, []int{1})

	for i := 1; i < numRows; i++ {
		//fmt.Println("ans:", ans)
		temp := make([]int, i+1)
		temp[0] = 1
		temp[len(temp)-1] = 1
		for j := 1; j < i; j++ {
			temp[j] = ans[i-1][j] + ans[i-1][j-1]
		}
		//fmt.Println("temp:", temp)
		ans = append(ans, temp)
	}
	return ans
}

func main() {
	ans := generate(10)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
