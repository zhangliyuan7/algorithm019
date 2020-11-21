package Week_02

import (
	"fmt"
	"testing"
)

//func TestBHeap(t *testing.T) {
//	var b bHeap
//	b.Push(10)
//	b.Push(11)
//	b.Push(12)
//	b.Push(22)
//	b.Push(9)
//	b.Push(7)
//	b.Push(13)
//	b.Push(21)
//	b.Push(32)
//	b.Tree()
//	fmt.Println(b.Top())
//	b.Pop()
//	b.Pop()
//	fmt.Println(b.Top())
//	b.Tree()
//	for i:=0;i<3;i++{
//		fmt.Println(b.Top())
//		b.Pop()
//	}
//	b.Push(999)
//	b.Tree()
//	fmt.Println(b.Top())
//}

func TestBHeap2(t *testing.T) {
	var b bheap
	b.Push(1)
	b.Push(3)
	b.Push(5)
	b.Push(2)
	b.Push(0)
	b.Push(77)
	b.Push(9)
	fmt.Println(b.h)
	fmt.Println(b.TOP())
	fmt.Println(b.Pop())
	fmt.Println(b.TOP())
	fmt.Println(b.Pop())
	b.Push(13)
	fmt.Println(b.h)
	fmt.Println(b.Pop())
	fmt.Println(b.Pop())
	fmt.Println(b.Pop())
}