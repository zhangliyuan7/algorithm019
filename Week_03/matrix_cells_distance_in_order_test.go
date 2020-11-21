package w3

import (
	"fmt"
	"testing"
)

func TestAllCellsDistOrder1(t *testing.T) {
	var a = make(map[[2]int]bool)
	var r = [2]int{0,1}
	var s = [2]int{0,1}
	var z = [2]int{1,0}
	a[r]=true
	if v,ok:=a[s];ok{
		fmt.Println(v)
	}
	if v,ok:=a[z];ok{
		fmt.Println(v)
	}
}