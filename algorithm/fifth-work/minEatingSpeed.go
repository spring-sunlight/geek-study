package fifth_work

import (
	"math"
	"sort"
)

//利用了 go 的官方库

func minEatingSpeed(piles []int, h int) int {
	var max int
	for i, pile := range piles {
		if i == 0 || pile > max {
			max = pile
		}
	}
	return sort.Search(max, func(i int) bool {
		need := 0
		for _, pile := range piles {
			need += int(math.Ceil(float64(pile) / float64(i+1)))
		}
		return need <= h
	}) + 1
}
