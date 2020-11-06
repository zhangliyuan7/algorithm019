package practice

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var list = []int{1,2,3}
	var head = &ListNode{}
	var firstNode = head
	for _,v:=range list{
		head.Val=v
		head.Next=&ListNode{}
		head=head.Next
	}
	SwapPairs(firstNode)
}