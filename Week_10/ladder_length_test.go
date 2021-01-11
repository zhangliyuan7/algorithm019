package Week_10

import (
	"fmt"
	"testing"
)

func TestLadderLength2(t *testing.T){
	fmt.Println(LadderLength("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
}
