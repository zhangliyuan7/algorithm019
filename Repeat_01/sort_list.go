package Repeat_01

//148
// 归并
func sortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	f:=head.Next
	s:=head
	for f!=nil&&f.Next!=nil{
		f=f.Next.Next
		s=s.Next
	}
	rnode:=s.Next
	lnode:=head
	s.Next=nil
	return merge(sortList(lnode),sortList(rnode))
}

func merge(l1,l2 *ListNode)*ListNode{
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val>l2.Val{
		l2.Next=merge(l1,l2.Next)
		return l2
	}else{
		l1.Next=merge(l1.Next,l2)
		return l1
	}
}