package w3

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	var gas=[]int{3,1,1}
	var cost=[]int{1,2,2}
	fmt.Println(CanCompleteCircuit(gas,cost))
}