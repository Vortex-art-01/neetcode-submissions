type MinStack struct {
	stack [][]int
}

func Constructor() MinStack {
	return MinStack{
		stack: [][]int{},
	}
}

func (this *MinStack) Push(val int) {

	min := val
	if len(this.stack) > 0 {
		min2 := this.GetMin()
		if min2 < min {
			min = min2
		}
	}

	this.stack = append(this.stack, []int{val, min})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}
