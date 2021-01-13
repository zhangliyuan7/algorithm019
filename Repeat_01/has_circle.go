package Repeat_01

// 141
func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	fast:=head.Next.Next
	slow:=head.Next
	for slow!=nil&&fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			return true
		}
	}
	return false
}