package Repeat_01

import (
	"sort"
	"strconv"
)

func Interview(s string )map[int]int{
	left:=0
	right:=0
	ln:=len(s)
	m:=make(map[int]int)
	for i:=0;i<ln;i++{
		if s[i]<='0'||s[i]>='9'{
			if left!=right{
				k,_:=strconv.Atoi(s[left:right])
				m[k]+=1
				left=right
			}
			left++
			right++
		}else{
			right++
		}
	}
	if right!=left{
		k,_:=strconv.Atoi(s[left:right])
		m[k]+=1
	}
	sort.Sort()
	return m
}
