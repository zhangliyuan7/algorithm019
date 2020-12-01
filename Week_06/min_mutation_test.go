package Week_06

import (
	"fmt"
	"testing"
)

func TestMinMutation(t *testing.T) {
	fmt.Println(MinMutation("AACCGGTT","AAACGGTA", []string{"AACCGGTA","AACCGCTA","AAACGGTA"}))
}