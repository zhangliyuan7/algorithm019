package w3

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}


func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	lnode := sortList1(head)
	rnode := sortList1(mid)
	return merge2SortedListNode(lnode, rnode)
}

func merge2SortedListNode(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = merge2SortedListNode(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge2SortedListNode(list1, list2.Next)
		return list2
	}
}
//buhao
//空间复杂度太高
//时间复杂度
func sortList2(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	dict:=make(map[int][]*ListNode)
	sortkeys:=[]int{}
	for head!=nil{
		dict[head.Val]=append(dict[head.Val],head)
		sortkeys=append(sortkeys,head.Val)
		head=head.Next
	}
	sort.Ints(sortkeys)
	var fir = &ListNode{}
	a:=fir
	for i:=0;i<len(sortkeys);i++{
		for _,v:=range dict[sortkeys[i]]{
			a.Next=v
			a=a.Next
		}
	}
	return fir.Next
}

// pass 可以一战
func sortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	var a = head
	for a.Next!=nil{
		a=a.Next
	}
	return sortX(head,a)
}

func sortX(head,tail *ListNode)*ListNode{
	half:=&ListNode{}
	if head==tail{
		head.Next=nil
		return head
	}
	if head.Next==tail{
		if head.Val>tail.Val{
			tail.Next=head
			head=tail
			tail=tail.Next
		}
		tail.Next=nil
		return head
	}
	var f, s = head, head
	for f!=tail&&f.Next!=tail&&f.Next.Next!=tail{
		f = f.Next.Next
		s = s.Next
	}
	half=s
	halfNext:=s.Next
	return mergeSortedLinkedList(sortX(head,half),sortX(halfNext,tail))
}

func mergeSortedLinkedList(link1,link2 *ListNode)*ListNode{
	var r =&ListNode{}
	a:=r
	for link1!=nil&&link2!=nil{
		if link1.Val>link2.Val{
			a.Next=link2
			link2=link2.Next
			a=a.Next
		}else{
			a.Next=link1
			link1=link1.Next
			a=a.Next
		}
	}
	if link1==nil{
		a.Next=link2
		return r.Next
	}
	a.Next=link1
	return r.Next
}

