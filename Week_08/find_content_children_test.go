package Week_08

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	fmt.Println(FindContentChildren([]int{1,2,4},[]int{1,3,1,1}))
}