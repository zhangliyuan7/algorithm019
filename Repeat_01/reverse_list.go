package Repeat_01

// 206
func reverseList(head *ListNode) *ListNode {
	j:=new(ListNode)
	for head!=nil{
		next:=j.Next
		j.Next=head
		head=head.Next
		j.Next.Next=next
	}
	return j.Next
}