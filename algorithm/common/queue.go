package common

import "fmt"

type queue struct {
	nums []interface{}
}

func NewQueue() *queue {
	return &queue{
		nums: []interface{}{},
	}
}

func (q *queue) IsEmpty() bool {
	return len(q.nums) == 0
}

func (q *queue) First() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.nums[0]
}

func (q *queue) Last() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.nums[len(q.nums)-1]
}

func (q *queue) Push(val interface{}) {
	q.nums = append(q.nums, val)
}

func (q *queue) Pop() {
	if !q.IsEmpty() {
		q.nums = q.nums[1:]
	}
}

func (q *queue) PrintQueue() {
	fmt.Println("==================")
	fmt.Println("queue 的长度为:", len(q.nums))
	fmt.Println("queue 的值为:", q.nums)
}
