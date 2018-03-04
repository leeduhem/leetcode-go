package p155

type MinStack struct {
	min     int
	stack   []int
	isEmpty bool
}

func Constructor() MinStack {
	return MinStack{
		isEmpty: true,
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if this.isEmpty || x < this.min {
		this.min = x
	}
	this.isEmpty = false
}

func (this *MinStack) Pop()  {
	l := len(this.stack)
	x := this.stack[l-1]
	this.stack = this.stack[:l-1]
	if l == 1 {
		this.isEmpty = true
	} else if x == this.min {
		this.min = this.stack[0]
		for _, e := range this.stack {
			if e <= this.min {
				this.min = e
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
