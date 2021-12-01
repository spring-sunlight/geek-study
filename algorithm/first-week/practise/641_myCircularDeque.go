package practise

import "fmt"

type MyCircularDeque struct {
	size  int
	deque []int
}

func ConstructorMyCircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		size:  k,
		deque: []int{},
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.deque = append([]int{value}, this.deque...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.deque = append(this.deque, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.deque = this.deque[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.deque = this.deque[:len(this.deque)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[len(this.deque)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.deque) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.deque) == this.size
}

func TestMyCircularDeque() {
	obj := ConstructorMyCircularDeque(3)
	fmt.Println(obj.InsertLast(1))
	fmt.Println(obj.InsertLast(2))
	fmt.Println(obj.InsertFront(3))
	fmt.Println(obj.InsertFront(4))
	fmt.Println(obj.GetRear())
	fmt.Println(obj.IsFull())
	fmt.Println(obj.DeleteLast())
	fmt.Println(obj.InsertFront(4))
	fmt.Println(obj.GetFront())

}
