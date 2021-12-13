package example_twoPointers_slidingWindow

import "fmt"

func twoSumII(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			return []int{left, right}
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSumII(numbers, 9))
}
