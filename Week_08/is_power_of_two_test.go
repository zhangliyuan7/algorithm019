package Week_08

import (
	"fmt"
	"testing"
)

//231

func TestA(t *testing.T){
	n:=15
	fmt.Println(n&(-n))
}
func TestB(t *testing.T){
	n:=1011111
	fmt.Println(n<<1)
}
func TestC(t *testing.T){
	n:=uint32(11001)
	fmt.Println(uint32(n>>3))
}