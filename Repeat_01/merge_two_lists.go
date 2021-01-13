package Repeat_01

// 21

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r:=new(ListNode)
	rgo:=r
	for ;l1!=nil&&l2!=nil;r=r.Next{
		if l1.Val>l2.Val{
			r.Next=l2
			l2=l2.Next
			continue
		}
		r.Next=l1
		l1=l1.Next
	}
	if l1!=nil{
		r.Next=l1
	}
	if l2!=nil{
		r.Next=l2
	}
	return rgo.Next
}