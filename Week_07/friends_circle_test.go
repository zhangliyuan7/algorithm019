package Week_07

import (
	"fmt"
	"testing"
)

func TestFindCircleNum(t *testing.T){
	fmt.Println(FindCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}}))
}