package Repeat_01

import (
	"fmt"
	"testing"
	"time"
)

//func TestMedianSlidingWindow(t *testing.T){
//	fmt.Println(medianSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
//}


//func TestMedianSlidingWindowReDo(t *testing.T){
//	fmt.Println(medianSlidingWindowReDo([]int{1,3,-1,-3,5,3,6,7},4))
//}
//3
//-1,1 => -3

func TestTimeParse(t *testing.T){
	tm:=time.Unix(1612354364,0)

	//v,err:=time.Parse("2006-01-02 15",	tm.String())

	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
}