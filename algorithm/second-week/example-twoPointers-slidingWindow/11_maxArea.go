package example_twoPointers_slidingWindow

import "fmt"

func maxArea(height []int) int {
	ans := 0
	for left, right := 0, len(height)-1; left < right; {
		m := min(height[left], height[right])
		if m*(right-left) > ans {
			ans = m * (right - left)
		}
		if m == height[left] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
