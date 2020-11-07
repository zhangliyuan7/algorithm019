package practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
//三个节点一个操作单位，所以能保持连接关系，操作后两个节点就行了
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var firNode = head.Next // return the second node
	for head != nil && head.Next.Next != nil {
		tempNode := head.Next.Next //temp = 3
		if tempNode != nil && tempNode.Next != nil {
			head.Next = tempNode.Next // 1.next= 4
		} else {
			head.Next = tempNode // 4==nil => 1.next=3
		}
		head.Next.Next = head //2.next=1
		head = tempNode      //temp is 3 ,下一次处理3
	}
	return firNode
}

func SwapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummy.Next = head
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next
		temp.Next = end
		start.Next = end.Next
		end.Next = start
		temp = start
	}
	return dummy.Next
}

// zero , three nodes join in once loop
// zero.next = b
// a.next=b.next
// b.next=a
// zero= a
//return zero.next
func SwapPairs3(head *ListNode) *ListNode {
	var pre =&ListNode{Val: 0,Next: head}
	var tmp = pre
	for tmp!=nil&&tmp.Next!=nil{
		start:=tmp.Next
		end:=tmp.Next.Next
		tmp.Next=end
		start.Next=end.Next
		end.Next=start
		tmp=start
	}
	return pre.Next
}

func SwapPairs4(head *ListNode) *ListNode {
	var pre = &ListNode{0,head}
	var tmp = pre
	for tmp!=nil{
		start:=tmp.Next
		end:=tmp.Next.Next
		start.Next=end.Next
		end.Next=start
		tmp.Next=end
		tmp=start
	}
	return pre.Next
}


