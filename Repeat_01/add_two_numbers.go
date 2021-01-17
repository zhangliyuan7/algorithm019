package Repeat_01

// 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一定要是指针 nil 方便判断head 是否存在，这个地方很巧妙
	var head *ListNode
	var tail *ListNode
	carry:=0

	for l1!=nil||l2!=nil{
		n1,n2:=0,0
		if l1!=nil{
			n1=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			n2=l2.Val
			l2=l2.Next
		}
		sum:=n1+n2+carry
		sum,carry=sum%10,sum/10
		//下面的if else比较巧妙，原地走两次，向后推进
		if head==nil{
			head=&ListNode{Val: sum}
			tail=head
		}else{
			tail.Next=&ListNode{Val: sum}
			tail=tail.Next
		}
	}
	if carry>0{
		tail.Next=&ListNode{Val: carry}
	}
	return head
}