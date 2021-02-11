package Repeat_01

import (
	"container/heap"
	"fmt"
	"sort"
)

//703
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

//调用heap接口只需要append进去即可 默认小顶堆
func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

//调用heap接口 只需要将切片长度减一即可
func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	fmt.Println("add before",kl.IntSlice)

	heap.Push(kl, val)
	fmt.Println("add end",kl.IntSlice)

	if kl.Len() > kl.k {
		fmt.Println("pop before",kl.IntSlice)

		heap.Pop(kl)
		fmt.Println("pop end",kl.IntSlice)

	}
	return kl.IntSlice[0]
}


//
//type KthLargest struct {
//	smallTopKHeap Heap //大顶堆 大小为k
//	largeTopHeap Heap //大顶堆 大小为n-k
//	k int
//}
//type Heap struct{
//	s []int
//}
//
//func (h *Heap)top()int {
//	return h.s[0]
//}
//
//func (h *Heap)push(x int ){
//	h.s=append(h.s,x)
//	//fmt.Println("before up",h.s)
//	h.heapifyup()
//	//fmt.Println("end up",h.s)
//
//}
//
//func (h *Heap)pop()int {
//	x:=h.s[0]
//	h.s[0]=h.s[len(h.s)-1]
//	h.s=h.s[:len(h.s)-1]
//
//	h.heapifydown()
//	return x
//}
//
//func (h *Heap)heapifydown(){
//	fa:=0
//	ch1,ch2:=1,2
//	for ch1<len(h.s)&&ch2<len(h.s){
//		if h.s[ch1]>h.s[ch2]{
//			h.swap(fa,ch1)
//			fa=ch1
//		}else{
//			h.swap(fa,ch2)
//			fa=ch2
//		}
//		ch1,ch2=fa*2+1,fa*2+2
//	}
//}
//
//func (h *Heap)heapifyup(){
//	last:=len(h.s)-1
//	fa:=(last-1)/2
//	for h.s[last]>h.s[fa]{
//		h.swap(fa,last)
//		last=fa
//		fa=(last-1)/2
//	}
//}
//
//func (h *Heap)swap(i,j int ){
//	h.s[i],h.s[j]=h.s[j],h.s[i]
//}
//
//func newHeap()Heap{
//	return Heap{[]int{}}
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	x:=KthLargest{
//		smallTopKHeap:newHeap(),
//		largeTopHeap:newHeap(),
//		k:k,
//	}
//	for i:=range nums{
//		x.smallTopKHeap.push(-nums[i])
//	}
//	x.sliding()
//	return x
//}
//
//func (this *KthLargest) Add(val int) int {
//	if val> -this.smallTopKHeap.top(){
//		this.smallTopKHeap.push(-val)
//		x:=this.smallTopKHeap.pop()
//		this.largeTopHeap.push(-x)
//		return -this.smallTopKHeap.top()
//	}
//	this.largeTopHeap.push(val)
//	return -this.smallTopKHeap.top()
//
//}
//
//func (this *KthLargest)sliding(){
//	if len(this.smallTopKHeap.s)<this.k{
//		this.smallTopKHeap.push(-this.largeTopHeap.pop())
//	}
//	if len(this.smallTopKHeap.s)>this.k{
//		this.largeTopHeap.push(-this.smallTopKHeap.pop())
//	}
//}
//
//
///**
// * Your KthLargest object will be instantiated and called as such:
// * obj := Constructor(k, nums);
// * param_1 := obj.Add(val);
// */