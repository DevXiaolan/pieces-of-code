package main

import (
	"fmt"
)

type MaxQueue struct {
	data []int
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: []int{},
		max:  []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) > 0 {
		return this.max[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)

	for i := len(this.max) - 1; i >= 0; i-- {
		if this.max[i] < value {
			this.max = this.max[0:i]
		} else {
			break
		}
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	target := this.data[0]
	this.data = this.data[1:]
	if len(this.max) > 0 && target == this.max[0] {
		this.max = this.max[1:]
	}
	return target
}

func main() {
	q := Constructor()
	q.Push_back(243)
	fmt.Printf("%+v\n", []interface{}{q})

}
