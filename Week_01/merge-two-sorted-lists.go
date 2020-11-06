package practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//bad case
//todo review
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs = &ListNode{}
	var tmp = rs
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else {
			tmp.Next = l1
			l1 = l1.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	}
	if l2 == nil {
		tmp.Next = l1
	}
	return rs.Next
}

//success
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs =&ListNode{}
	var tmp=rs
	for l1!=nil && l2!=nil{
		if  l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		}
		if l1.Val<l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		}
		if l1.Val==l2.Val{
			tmp.Next=l1
			l1=l1.Next
			tmp=tmp.Next
			tmp.Next=l2
			l2=l2.Next
		}
		tmp=tmp.Next
	}
	if l1==nil{
		tmp.Next=l2
	}
	if l2==nil{
		tmp.Next=l1
	}
	return rs.Next
}