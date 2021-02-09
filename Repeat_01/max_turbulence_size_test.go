package Repeat_01

import (
	"fmt"
	"testing"
)

func TestMaxTurbulenceSize(t *testing.T){
	//fmt.Println(maxTurbulenceSize([]int{9,4,2,10,7,8,8,1,9}))
	//fmt.Println(maxTurbulenceSizeSlidingWindow([]int{9,4,2,10,7,8,8,1,9}))
	fmt.Println(maxTurbulenceSizeSlidingWindow([]int{4,8,12,16,1}))
}
