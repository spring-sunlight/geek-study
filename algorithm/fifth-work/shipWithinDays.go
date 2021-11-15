package fifth_work

import "sort"

//利用了 go 的官方库

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= days
	})
}
