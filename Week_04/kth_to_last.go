package Week_04


//offer 22
type ListNode struct {
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head==nil{
		return nil
	}
	var a int =1
	r:=head
	for head.Next!=nil{
		if a==k{
			r=r.Next
		}
		if a<k{
			a+=1
		}
		head=head.Next
	}
	return r
}