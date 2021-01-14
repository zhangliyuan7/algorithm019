package Repeat_01

// offer 22
func getKthFromEnd(head *ListNode, k int) *ListNode {
	f:=head
	s:=head
	for f!=nil{
		f=f.Next
		k--
		if k<0{
			s=s.Next
		}
	}
	return s
}
