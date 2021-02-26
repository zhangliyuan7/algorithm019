package Repeat_01

import (
	"fmt"
	"testing"
)

func TestFindNumberOfValidWord(t *testing.T){
	words:=[]string{"aaaa","asas","able","ability","actt","actor","access"}
	puzzles:=[]string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}
	fmt.Println(findNumOfValidWords(words,puzzles))
}
