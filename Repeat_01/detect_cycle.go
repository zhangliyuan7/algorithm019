package Repeat_01

// 142
// a,b,c
// 1 必定在slow的第一圈相遇
// 2 a+b=slow 路程 a+n(b+c)+b是快的路程
// 快的是慢的二倍 so  a+n(b+c)+b=2(a+b)
//	=> a+(n+1)b+nc=2a+2b
//  => (n-1)b+nc=a => (n-1)(b+c)+c=a  又b+c等于环周长，故
//  当从快慢指针相遇的此时，从起点出发的步长为1的指针
//  将与慢指针将相遇于入环点 （慢指针刚好能绕n-1圈加上c的长度，走到入环口）

func detectCycle(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return nil
	}
	fast:=head
	slow:=head
	ptr:=head
	for fast!=nil{
		if fast.Next==nil{
			return nil
		}
		fast=fast.Next.Next
		slow=slow.Next
		if slow==fast{
			for ptr!=slow{
				ptr=ptr.Next
				slow=slow.Next
			}
			return slow
		}
	}
	return nil
}