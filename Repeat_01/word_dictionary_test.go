package Repeat_01

import (
	"fmt"
	"testing"
)

func TestWordDictionary_Search(t *testing.T) {
	wd:=Constructor()
	wd.AddWord("bad")
	fmt.Println(wd.next)
	fmt.Println(wd.Search("b.."))
}