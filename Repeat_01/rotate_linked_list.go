package Repeat_01

//61
func rotateRightI(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	f := head
	tmp := head
	l := 0
	end := new(ListNode)
	for tmp != nil {
		end = tmp
		l++
		tmp = tmp.Next
	}
	k = k % l
	if k == 0 {
		return head
	}
	last := new(ListNode)
	for ; l-k > 0; l-- {
		last = f
		f = f.Next
	}
	// 走了l-k步
	last.Next = nil
	end.Next = head
	return f
}