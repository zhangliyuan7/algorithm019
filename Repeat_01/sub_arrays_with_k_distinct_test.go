package Repeat_01

import (
	"fmt"
	"testing"
)

func TestSubArraysWithKDistinct(t *testing.T){
	fmt.Println(subarraysWithKDistinctSlidingWindow([]int{1,2,1,2,3},2))
}