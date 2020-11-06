package practice

import (
	"fmt"
	"testing"
)

func Test_MoveZero(t *testing.T) {
	var nums = []int{0,0,1}
	MoveZero(nums)
	fmt.Print(nums)
}
func Test_MoveZero2(t *testing.T) {
	var nums = []int{0,0,1}
	MoveZero2(nums)
	fmt.Print(nums)
}