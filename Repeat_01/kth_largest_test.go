package Repeat_01

import (
	"fmt"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	k:=Constructor(4,[]int{4,5,8,2})

	fmt.Println(k.Add(3))

	fmt.Println(k.Add(5))

	fmt.Println(k.Add(10))

	fmt.Println(k.Add(9))

	fmt.Println(k.Add(4))

}
