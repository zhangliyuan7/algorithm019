package Week_10

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T){
	nums:=[]int{1,2,3,4,5,6,7}
	Rotate(nums,1)
	//Roll(nums)
	//Roll(nums)
	//Roll(nums)

	fmt.Println(nums)

}
func TestReverse(t *testing.T){
	//nums:=[]int{1,2,3,4,5,6,7}
	nums:=[]int{1,2,3,4,5,6}
	Reverse(nums)
	fmt.Println(nums)
}

func TestRotateSmart(t *testing.T){
	nums:=[]int{1,2,3,4,5,6,7}
	RotateSmart(nums,3)
	fmt.Println(nums)
}