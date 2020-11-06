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