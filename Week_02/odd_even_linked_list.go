package Week_02

type ListNode struct {
	Val  int
	Next *ListNode
}
//总分总 最后合并
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenhead := head.Next
	p1 := head
	p2 := evenhead
	for p1.Next != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = evenhead
	return head
}
