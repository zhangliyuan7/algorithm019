package Repeat_01

//341
// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}


type NestedIterator struct {
	stack [][]*NestedInteger
}

func ConstructorNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}}

func (this *NestedIterator) Next() int {
	queue:=this.stack[len(this.stack)-1]
	val:=queue[0].GetInteger()
	this.stack[len(this.stack)-1] = queue[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		queue := this.stack[len(this.stack)-1]
		if len(queue) == 0 { // 当前队列为空，出栈
			this.stack = this.stack[:len(this.stack)-1]
					continue
		}
		nest:=queue[0]
		if nest.IsInteger(){
			return true
		}
		this.stack[len(this.stack)-1]=queue[1:]
		this.stack=append(this.stack,nest.GetList())
	}
	return false 
}
