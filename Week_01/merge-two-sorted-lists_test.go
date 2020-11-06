package practice

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	var l1l = []int{1,2,4}
	var l2l= []int{1,3,4}
	var l1 =&ListNode{}
	var l2 =&ListNode{}
	tmpl1:=l1
	tmpl2:=l2
	for i:=range l1l{
		tmpl1.Next=&ListNode{Val: l1l[i],Next: nil}
		tmpl1=tmpl1.Next
	}
	for i:=range l2l{
		tmpl2.Next=&ListNode{Val: l2l[i],Next: nil}
		tmpl2=tmpl2.Next
	}
	res:=MergeTwoLists(l1.Next,l2.Next)
	for res!=nil{
		fmt.Print(res.Val)
		res=res.Next
	}

}