package Repeat_01

import (
	"fmt"
	"testing"
)
var strs =[]string{"tars","rats","arts","star"}

//func TestSimilarFunc(t *testing.T){
//	strMap:=make(map[string]int)
//	for i:=range strs{
//		strMap[strs[i]]=i
//	}
//	fmt.Println("map:",strMap)
//	fmt.Println(SearchSimilarWord("star",strMap))
//}
func TestNumSimilarGroups(t *testing.T) {
	fmt.Println(numSimilarGroups(strs))
}
