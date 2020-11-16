package w3

import (
	"fmt"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	a:=[][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}
	fmt.Println(ReconstructQueue(a))
}