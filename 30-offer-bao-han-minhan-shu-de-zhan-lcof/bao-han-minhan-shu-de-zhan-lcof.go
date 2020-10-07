package main

import "fmt"

// 辅助栈

type MinStack struct {
	data    []int
	minData []int
	index   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
		-1,
	}
}

func (this *MinStack) Push(x int) {
	if this.index == -1 {
		this.data = append(this.data, x)
		this.minData = append(this.minData, x)
		this.index++
		return
	}
	this.data = append(this.data, x)
	if x < this.minData[this.index] {
		this.minData = append(this.minData, x)
	} else {
		this.minData = append(this.minData, this.minData[this.index])
	}
	this.index++
}

func (this *MinStack) Pop() {
	if this.index == -1 {
		return
	}
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
	this.index--
}

func (this *MinStack) Top() int {
	return this.data[this.index]
}

func (this *MinStack) Min() int {
	return this.minData[this.index]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.minData)
	fmt.Println(obj.Min())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Min())
}
