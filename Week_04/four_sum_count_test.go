package Week_04

import (
	"fmt"
	"testing"
)

func TestFourSumCount(t *testing.T){
	A,B,C,D:=[]int{-1,-1},[]int{-1,1},[]int{-1,1},[]int{1,-1}
	fmt.Println(FourSumCount(A,B,C,D))
}