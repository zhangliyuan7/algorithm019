package Repeat_01


// 232 double stack = q
type MyQueue struct {
	s_top []int
	s_back []int
}

type stack []int

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{[]int{},[]int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.s_top=append(this.s_top, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ln:=len(this.s_top)
	for len(this.s_top)!=0{
		this.s_back=append(this.s_back,this.s_top[len(this.s_top)-1])
		this.s_top=this.s_top[:len(this.s_top)-1]
	}
	r:= this.s_back[ln-1]
	this.s_back=this.s_back[:ln-1]
	for len(this.s_back)!=0{
		this.s_top=append(this.s_top,this.s_back[len(this.s_back)-1])
		this.s_back=this.s_back[:len(this.s_back)-1]
	}
	return r
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s_top)>0{
		return this.s_top[0]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s_top)==0&&len(this.s_back)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

