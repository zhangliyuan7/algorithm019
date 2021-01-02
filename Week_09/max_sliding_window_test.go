package Week_09

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T){
	nums:=[]int{1,3,-1,-3,5,3,6,7}
	fmt.Println(MaxSlidingWindow(nums,3))
}