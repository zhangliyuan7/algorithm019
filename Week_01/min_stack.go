package practice

//type MinStack struct {
//	minList *Stack
//	normalList *Stack
//	len int
//}
//
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return  MinStack{
//		NewStack(),
//		NewStack(),
//		0,
//	}
//}
//
//
//func (this *MinStack) Push(x int)  {
//	if this.len==0{
//		this.minList.Push(x)
//		this.normalList.Push(x)
//		this.len++
//		return
//	}
//		if this.minList.Top().(int)>=x{
//			this.minList.Push(x)
//		}else{
//			this.minList.Push(this.minList.Top())
//		}
//	this.normalList.Push(x)
//	this.len++
//}
//
//
//func (this *MinStack) Pop()  {
//	this.minList.Pop()
//	this.normalList.Pop()
//	this.len--
//}
//
//
//func (this *MinStack) Top() int {
//	return this.normalList.Top().(int)
//}
//
//
//func (this *MinStack) GetMin() int {
//	 return this.minList.Top().(int )
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// use list
type MinStack struct {
	minList    []int
	normalList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.normalList) == 0 {
		this.normalList = append(this.normalList, x)
		this.minList = append(this.minList, x)
		return
	}
	if this.minList[len(this.minList)-1] <= x {
		this.minList = append(this.minList, this.minList[len(this.minList)-1])
	}
	if this.minList[len(this.minList)-1] > x {
		this.minList = append(this.minList, x)
	}
	this.normalList = append(this.normalList, x)
}

func (this *MinStack) Pop() {
	if len(this.normalList) == 0 {
		return
	}
	this.minList = this.minList[:len(this.minList)-1]
	this.normalList = this.normalList[:len(this.normalList)-1]
}

func (this *MinStack) Top() int {
	if len(this.normalList) == 0 {
		return -1
	}
	return this.normalList[len(this.normalList)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minList) == 0 {
		return -1
	}
	return this.minList[len(this.minList)-1]
}
