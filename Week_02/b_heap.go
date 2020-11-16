package w2

import "fmt"

type bHeap struct{
	tree []int
	size int
}
func (bh *bHeap)Tree(){
	fmt.Println(bh.tree)
}
func (bh *bHeap)Top()int{
	return bh.tree[0]
}

func (bh *bHeap)Pop(){
	last:=bh.size-1
	bh.tree[last],bh.tree[0]=bh.tree[0],bh.tree[last]
	bh.size-=1
	bh.tree=bh.tree[:last]
	bh.HeapifyDown()
}

func (bh *bHeap)Push(n int ){
	bh.tree=append(bh.tree,n)
	bh.size+=1
	bh.HeapifyUp()
}

func (bh *bHeap)HeapifyUp(){
	last:=bh.size-1
	for last!=(last-1)/2{
		if bh.tree[last] > bh.tree[(last-1)/2] {
			bh.tree[last],bh.tree[(last-1)/2]=bh.tree[(last-1)/2],bh.tree[last]
			last=(last-1)/2
		}
		if bh.tree[last]<bh.tree[(last-1)/2]{
			break
		}
	}
}

func (bh *bHeap)HeapifyDown(){
	i:=0
	for i*2+2<bh.size{
		biggerChild:=bh.max(i*2+1,i*2+2)
		bh.tree[biggerChild],bh.tree[i]=bh.tree[i],bh.tree[biggerChild]
		i=biggerChild
	}
}

func (bh *bHeap)max(a,b int )int{
	if bh.tree[a]>bh.tree[b]{
		return a
	}
	return b
}



type  bheap struct{
	h []int
}

func (h *bheap)TOP()int{
	return h.h[0]
}

func (h *bheap)Push(n int ){
	h.h=append(h.h, n)
	h.heapifyUp()
}

func (h *bheap)Pop()int{
	popN:=h.h[0]
	last:=h.h[len(h.h)-1]
	h.h[0]=last
	h.h=h.h[:len(h.h)-1]
	h.heapifyDown()
	return popN
}

func (h *bheap)heapifyUp(){
	lastId :=len(h.h)-1
	fatherId:=(lastId-1)/2
	for h.h[lastId]>h.h[fatherId] {
		swap(fatherId, lastId, h.h)
		lastId=fatherId
		fatherId=(lastId-1)/2
	}
}

func (h *bheap)heapifyDown(){
	var topId,childIdA,childIdB int
	childIdA=topId*2+1
	childIdB=topId*2+2
	for childIdA<len(h.h)&&childIdB<len(h.h){
		if h.h[childIdA]>h.h[childIdB]{
			swap(topId,childIdA,h.h)
			topId=childIdA
		}else{
			swap(topId,childIdB,h.h)
			topId=childIdB
		}
		childIdA=topId*2+1
		childIdB=topId*2+2
	}
}

func swap(a,b int ,list []int){
	list[a],list[b]=list[b],list[a]
}