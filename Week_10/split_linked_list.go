package Week_10

// 86
type ListNode struct {
    Val int
    Next *ListNode
}
func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	tmpLess:=less
	greater :=&ListNode{}
	tmpGreater:=greater
	for head!=nil{
		if head.Val<x{
			tmpLess.Next=head
			tmpLess=tmpLess.Next
		}else{
			tmpGreater.Next=head
			tmpGreater=tmpGreater.Next
		}
		tmp:=head.Next
		head.Next=nil
		head=tmp
	}
	tmpLess.Next=greater.Next
	return less.Next
}