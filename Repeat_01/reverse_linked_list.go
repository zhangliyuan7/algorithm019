package Repeat_01

// 92


func reverseBetweenI(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 保存的是整个链表当前反转部分的最后一个节点
	// 始终保持 cur乃反转部分的尾巴，cur的next存放非反转部分的首个
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next // 因为这个节点仍需要反转 ，将后一个节点提取出来
		cur.Next = next.Next // cur的下一个节点就保存反转尾巴后面的那个节点
		next.Next = pre.Next  // 反转操作
		pre.Next = next   // 反转操作
	}
	return dummyNode.Next
}

func reverseBetweenII(head *ListNode, left int, right int) *ListNode {
	// (left-1)th node
	before:=&ListNode{}
	// first node
	start:=before
	for i:=1;i<=right;i++{
		// reverse start from left node
		if i==left{
			// reverse-linkedlist start
			pre:=&ListNode{}
			// reverse-linkedlist end
			last:=head
			for ;left<=right;left++{
				tmp:=pre.Next
				pre.Next=head
				head=head.Next
				pre.Next.Next=tmp
			}
			// head is the (right+1)th node now
			// assembly(组装/部件)
			last.Next=head
			before.Next=pre.Next
			break
		}
		// keep record (left-1)th node
		before.Next=head
		before=before.Next
		// loop
		head=head.Next
	}
	return start.Next
}

func reverseBetweenIII(head *ListNode, m int, n int) *ListNode {
	x := new(ListNode)
	x.Next = head
	oldStart := x
	oldEnd := new(ListNode)
	reverseList := ListNode{}
	reverseEnd := &ListNode{}
	for x != nil {
		m--
		n--
		if m == 0 {
			oldEnd = x
		}
		tmp := x.Next
		if m <= 0 && n >= 0 {
			a := x.Next
			x.Next = a.Next
			a.Next = reverseList.Next
			reverseList.Next = a
			if m == 0 {
				reverseEnd = reverseList.Next
			}
		} else {
			x = tmp
		}

	}
	last := oldEnd.Next
	reverseEnd.Next = last
	oldEnd.Next = reverseList.Next
	return oldStart.Next
}
