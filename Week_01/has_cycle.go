package practice


func HasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil||head.Next.Next==nil{
		return false
	}
	slowP:=head.Next
	fastP:=head.Next.Next
	for fastP.Next!=nil&&fastP.Next.Next!=nil{
		if slowP==fastP{
			return true
		}
		slowP=slowP.Next
		fastP=fastP.Next.Next
	}
	return false
}

func HasCycle2(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	fastP :=head.Next.Next
	slowP:= head.Next
	for fastP!=nil&&fastP.Next!=nil{
		if fastP==slowP{
			return true
		}
		fastP=fastP.Next.Next
		slowP=slowP.Next
	}
	return false
}


func CycleHas ( n *ListNode)bool{
	if n.Next==nil||n.Next.Next==nil{
		return false
	}
	fast:=n.Next.Next
	slow:=n.Next
	for fast!=nil&&fast.Next!=nil{
		if fast == slow {
			return true
		}
		fast=fast.Next.Next
		slow=slow.Next
	}
	return false
}