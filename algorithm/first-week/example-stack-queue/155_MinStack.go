package example_stack_queue

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	val  int
	list *list.List
}

func ConstructorMinStack() MinStack {
	return MinStack{
		list: list.New(),
	}
}

func (this *MinStack) Push(val int) {
	this.list.PushBack(val)
}

func (this *MinStack) Pop() {
	fmt.Println(this.list.Back().Value.(int))
	this.list.Remove(this.list.Back())
}

func (this *MinStack) Top() int {
	return this.list.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	e := this.list.Front()
	minVal := e.Value.(int)
	for e != nil {
		if minVal > e.Value.(int) {
			minVal = e.Value.(int)
		}
		e = e.Next()

	}
	return minVal
}

func main() {
	obj := ConstructorMinStack()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
