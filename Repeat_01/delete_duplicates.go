package Repeat_01

//82
func deleteDuplicates(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	dummy:=&ListNode{}
	dummy.Next=head
	cur:=dummy
	for cur.Next!=nil&&cur.Next.Next!=nil{
		if cur.Next.Val==cur.Next.Next.Val{
			val:=cur.Next.Val
			for cur.Next!=nil&&cur.Next.Val==val{
				cur.Next=cur.Next.Next
			}
		}else{
			cur=cur.Next
		}
	}
	return dummy.Next
}
