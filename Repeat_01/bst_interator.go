package Repeat_01


//173

//
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	index int
	cap int
	serial []int
}


func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	// index:=0
	// cap:=calculateCap(root)
	node:=BSTIterator{}
	node.midOrder(root)
	node.cap=len(node.serial)
	return node
}


func (this *BSTIterator) Next() int {
	if this.index<this.cap{
		res:=this.serial[this.index]
		this.index+=1
		return res
	}
	return 0
}


func (this *BSTIterator) HasNext() bool {
	if this.index<this.cap{
		return true
	}
	return false
}

// func calculateCap(root *TreeNode)int{
//     cap:=0
//     if root!=nil{
//         cap+=1
//         cap+=calculateCap(root.Left)
//         cap+=calculateCap(root.Right)
//     }
//     return cap
// }

func (this *BSTIterator)midOrder(root *TreeNode){
	if root!=nil{
		this.midOrder(root.Left)
		this.serial=append(this.serial,root.Val)
		this.midOrder(root.Right)
	}
}


//smart way : use stack

type BSTIteratorI struct {
	 s []*TreeNode
	 cur *TreeNode
}


func ConstructorBSTIteratorI(root *TreeNode) BSTIteratorI {
	return BSTIteratorI{cur: root}
}


func (this *BSTIteratorI) Next() int {
	for node:=this.cur;node!=nil;node=node.Left{
		this.s=append(this.s,node)
	}
	this.cur,this.s=this.s[len(this.s)-1],this.s[:len(this.s)-1] //当前中序最左边的元素 ，pop后的栈
	val:=this.cur.Val
	this.cur=this.cur.Right //中序的保证
	return val
}


func (this *BSTIteratorI) HasNext() bool {
	if this.cur!=nil|| len(this.s)!=0{
		return true
	}
	return false
}
