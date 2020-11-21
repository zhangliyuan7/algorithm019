package w3

import (
	"sort"
)

// NO.406
func ReconstructQueue(people [][]int) [][]int {
	sort.Slice(people,func(i,j int )bool{
		if people[i][0]==people[j][0]{
			return people[i][1]<people[j][1]
		}
		return people[i][0]>people[j][0]
	})
	for i,v:=range people{
		copy(people[v[1]+1:i+1],people[v[1]:i])
		people[v[1]]=v
	}
	return people
}

func ReconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0]>people[j][0]||
			(people[i][0]==people[j][0]&&people[i][1]<people[j][1])
	})
	for i,v:=range people{
		indexV:=v[1]
		copy(people[indexV+1:i+1],people[indexV:i])
		people[v[1]]=v
	}
	return people
}