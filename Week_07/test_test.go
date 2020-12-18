package Week_07

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T){
	namesCandidate:=[]string{"Chandler","Joey","Max","Mike"}
	peopleCandidate:=[]string{"z-hua","fj","dbl","z-he"}
	names:=make(map[string]string)
	for i,_:=range namesCandidate{
		names[namesCandidate[i]]=""
	}
	index:=0
	for k,_:=range names{
		names[k]=peopleCandidate[index]
		index+=1
	}
	fmt.Println(names)
}