package practice

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head!=nil{
		next:= head.Next
		head.Next=pre
		pre=head
		head=next
	}
	return pre
}
// 类似栈过程 开一个新链表 往里从头塞
// 存A.next 也就是B
// 把pre给A.next
// A给pre
// B给head
//下次循环B开始 存C 然后pre给B.next B给pre，C给head

// use two pointers
//  new linked list
// save to new linked list by read sort
func ReverseLinkedList(head *ListNode)*ListNode{
	var pre = &ListNode{}
	for head!=nil{
		tail := head.Next
		head.Next=pre
		pre=head
		head=tail
	}
	return pre
}