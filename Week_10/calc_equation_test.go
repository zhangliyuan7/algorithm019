package Week_10

import (
	"fmt"
	"testing"
)

func TestCalcEquation(t *testing.T){
	fmt.Println(CalcEquation([][]string{{"a","b"},{"b","c"}},[]float64{2.0,3.0},[][]string{{"a","c"},{"b","a"},{"a","e"},{"a","a"},{"x","x"}}))
}
