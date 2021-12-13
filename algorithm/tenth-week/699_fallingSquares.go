package tenth_week

//借鉴了题解,朴素解法

func fallingSquares(positions [][]int) []int {
	/*
		n=1000，暴力处理即可，时间复杂度 O(n^2)
		维护当前所有不同高度的区间
		当新方块落下时，可能命中多个区间
			被部分命中的需要切割
			被完全命中的直接舍弃
		新方块的高度，则由被命中的最高区间决定
	*/

	n := len(positions)
	ret := make([]int, n)

	a := [][]int{}
	a = append(a, []int{0, int(1e9), 0})
	a, ret[0] = merge(a, positions[0])
	h := 0
	for i := 1; i < n; i++ {
		a, h = merge(a, positions[i])
		ret[i] = max(ret[i-1], h)
	}
	return ret
}

func merge(a [][]int, block []int) ([][]int, int) {
	left, width, right := block[0], block[1], block[0]+block[1]

	ret := [][]int{}
	i := 0
	// 保留左侧所有区间
	for i < len(a) && a[i][1] <= left {
		ret = append(ret, a[i])
		i++
	}
	// 至此，第i个区间的后半段被 block 命中，先保留前半段
	ret = append(ret, []int{a[i][0], left, a[i][2]})

	// 所有被命中的区间，取最高，并且被完全命中的，无需保留
	h := 0
	for i < len(a) && a[i][1] <= right {
		h = max(h, a[i][2])
		i++
	}
	// 至此，第i个区间的后半段超出了 block

	// 检查其前半段是否被命中
	if a[i][0] < right {
		h = max(h, a[i][2])
		ret = append(ret, []int{left, right, h + width})
		ret = append(ret, []int{right, a[i][1], a[i][2]}) // 切掉被命中的前半段，只保留后半段
		i++
	} else {
		ret = append(ret, []int{left, right, h + width})
	}

	// 保留右侧所有区间
	for i < len(a) {
		ret = append(ret, a[i])
		i++
	}
	return ret, h + width
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
