package forth_week

import "math"

//借鉴了题解

/*
解题思路
首先想到的是暴力破解：

1、最小的速度一小时吃一根，最快的速度一小时吃完最多的那一堆香蕉（再多了没有价值，因为吃完一堆也只能干看着）
2、从1到最大速度，依次看能否在规定时间吃完
3、第一个满足条件的速度，就是我们要的
但是，这样太慢了，我们可以用二分法来迅速找到这个速度

因为珂珂每次最多只能吃同一堆的香蕉，所以最大速度应该是最大的那堆香蕉，设定为right.再多了没有意义。
我们只要找到0到right中，满足条件的那个速度，让珂珂可以吃掉香蕉就好了。
很自然的想到了二分法。

4、left=1,right=最大速度，取其中间那个值mid
1）如果mid速度满足条件，说明速度还可以更小， right = mid
2）如果mid速度不满足条件，说明速度可以大一点，left = mid + 1
如何判断规定的速度可以吃完香蕉
对于每一堆香蕉，其耗费时间为，香蕉的量/速度的值向上取整。

计算总时间是否小于规定时间H就好了。

向上取整的特殊技巧
m / n 向上取整可以写成 (m + n - 1) / n
1）如果m/n可以整除，那么n-1/n小于1，会被略掉，不影响原来值
2）如果m/n不可以整除，那么剩余的那点点值，加上n-1/n，正好可以形成进位1，自然就是向上取整了
*/

func minEatingSpeed(piles []int, H int) int {
	left, right, mid := 1, 0, 0
	for _, v := range piles {
		if v > right {
			right = v
		}
	}
	for left < right {
		mid = (left + right) / 2
		if canDo(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 可以吃完所有香蕉
func canDo(piles []int, speed, H int) bool {
	time := 0
	for _, v := range piles {
		time += int(math.Ceil(float64(v) / float64(speed)))
	}
	return time <= H
}
