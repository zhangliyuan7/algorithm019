package Repeat_01

// 83
func deleteDuplicatesEasier(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	dummy:=&ListNode{}
	dummy.Next=head
	start:=dummy
	for head.Next!=nil{
		if head.Next.Val==head.Val{
			head.Next=head.Next.Next
		}else{
			head=head.Next
		}
	}
	return start.Next
}
