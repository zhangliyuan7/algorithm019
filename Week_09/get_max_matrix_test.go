package Week_09

import (
	"fmt"
	"testing"
)

func TestGetMaxMatrix(t *testing.T){
	fmt.Println(GetMaxMatrix([][]int{{-1,0},{0,-1}}))
}