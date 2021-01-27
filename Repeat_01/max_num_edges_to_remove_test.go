package Repeat_01

import (
	"fmt"
	"testing"
)

func TestMaxNumEdgesToRemove(t *testing.T){
	//edges:= [][]int{{3,2,3},{1,1,2},{2,3,4}}
	//edges:= [][]int{{3,1,2},{3,2,3},{1,1,4},{2,1,4}}
	edges:= [][]int{{3,1,2},{3,2,3},{1,1,3},{1,2,4},{1,1,2},{2,3,4}}
	fmt.Println(MaxNumEdgesToRemove(4, edges))
}