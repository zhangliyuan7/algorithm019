package Week_07

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T){
	fmt.Println(Exist([][]byte{[]byte{'A','B','C','E'},{'S','F','C','S'},{'A','D', 'E','E'}},"ABCB"))
}