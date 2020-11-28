package Week_04

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var r = &ListNode{}
	var res = r
	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			r.Next=l1
			l1=l1.Next
			r=r.Next
		}else{
			r.Next=l2
			l2=l2.Next
			r=r.Next
		}
	}
	if l1==nil{
		r.Next=l2
		return res.Next
	}
		r.Next=l1
		return res.Next
}