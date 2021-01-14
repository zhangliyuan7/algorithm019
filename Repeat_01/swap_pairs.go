package Repeat_01

// 24
func swapPairs(head *ListNode) *ListNode {
	m:=new(ListNode)
	r:=m
	for head!=nil&&head.Next!=nil{
		c:=head.Next.Next
		b:=head.Next
		a:=head
		m.Next=b
		b.Next=a
		a.Next=c
		// important
		m=a
		// important
		head=c
	}
	return r.Next
}