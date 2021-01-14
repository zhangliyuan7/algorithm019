package Repeat_01

// 92
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	x:=new(ListNode)
	x.Next=head
	oldStart:=x
	oldEnd:=new(ListNode)
	reverseList:=ListNode{}
	reverseEnd:=&ListNode{}
	for x!=nil{
		m--
		n--
		if m==0{
			oldEnd=x
		}
		tmp:=x.Next
		if m<=0&&n>=0{
			a:=x.Next
			x.Next=a.Next
			a.Next=reverseList.Next
			reverseList.Next=a
			if m==0{
				reverseEnd=reverseList.Next
			}
		}else{
			x=tmp
		}

	}
	last:=oldEnd.Next
	reverseEnd.Next=last
	oldEnd.Next=reverseList.Next
	return oldStart.Next
}