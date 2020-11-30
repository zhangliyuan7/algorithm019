package Week_04


func reverseList(head *ListNode) *ListNode {
	pre :=&ListNode{}
	for head!=nil{
		tail :=head.Next
		head.Next=pre.Next
		pre.Next=head
		head=tail
	}
	return pre.Next
}