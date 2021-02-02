package Repeat_01

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestCharacterReplacement(t *testing.T){
	sec, dec := math.Modf(1612255112.711511)
	fmt.Println(time.Unix(int64(sec), int64(dec*(1e9))))
	fmt.Println(time.Now().UnixNano())
	fmt.Println(characterReplacement("ABBB",2))
}
