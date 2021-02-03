package Repeat_01

import (
	"fmt"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T){
	fmt.Println(medianSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
}
