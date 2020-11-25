package Week_04

import (
	"sort"
	"strings"
)

func SortString2(s string)string{
	var r strings.Builder
    var sl [26]int
    for _,c:=range s{
    	sl[c-'a']+=1
	}
	for len(s)!=len(r.String()){
		for i:=0;i<26;i++{
			if sl[i]!=0{
				r.WriteByte(uint8(i)+'a')
				sl[i]-=1
			}
		}
		for i:=25;i>=0;i--{
			if sl[i]!=0{
				r.WriteByte(uint8(i)+'a')
				sl[i]-=1
			}
		}
	}
	return r.String()
}

//1370 左右扫描，代码有点太长了，不过效果还算不错
func SortString(s string) string {
	sb:=[]byte(s)
	sort.Slice(sb, func(i, j int) bool {
		if sb[i]<sb[j]{
			return true
		}
		return false
	})
	var r strings.Builder
	direct,i :=0,0
	tmp :=uint8(0)
	for len(r.String())!=len(s){
		if direct==0 {
			if i <= len(sb)-1 {
				if tmp == sb[i] {
					i++
					continue
				}
				r.WriteByte(sb[i])
				tmp = sb[i]
				sb = append(sb[0:i], sb[i+1:]...)
			} else {
				direct = 1
				i=len(sb)-1
				tmp=0
			}
		}else{
			if i >=0 {
				if tmp == sb[i] {
					i--
					continue
				}
				r.WriteByte(sb[i])
				tmp = sb[i]
				sb = append(sb[0:i], sb[i+1:]...)
				i--
			}else{
				direct=0
				i=0
				tmp=0
			}
		}
	}
	return r.String()
}