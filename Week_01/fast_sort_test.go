package practice

import (
	"fmt"
	"testing"
)

//func TestFastSortA(t *testing.T) {
//	var a = []int{1,4,5,11,6632,2,56,1,77,12}
//	FastSortA(a)
//	fmt.Println(a)
//}
func TestFastSortC(t *testing.T) {
	var a = []int{1,4,5,11,6632,2,56,1,77,12}
	Quick3Sort(a,0,len(a)-1)
	fmt.Println(a)
}