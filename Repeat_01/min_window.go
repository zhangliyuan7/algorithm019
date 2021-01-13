package Repeat_01

import "math"

// 76
// sliding window + hash table ?
func MinWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < len) {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
		//for mapCompare(keys,keysCompare)&&left<=right {
		//			if ansright==-1||right-left< ansright-ansleft{
		//				ansright=right+1
		//				ansleft=left
		//			}
		//			if v,ok:=keys[string(s[left])];ok{
		//				keysCompare[string(s[left])]=v-1
		//			}
		//			left++
		//		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func MinWindowB(s ,t string )string{
	keys:=make(map[string]int)
	for _,v:=range t{
		keys[string(v)]+=1
	}
	keysCompare:=make(map[string]int)
	ansleft:=-1
	ansright:=math.MaxInt32
	left:=0
	right:=0
	for ;right<len(s);right++{
		if v,ok:=keys[string(s[right])];ok&&v>0{
			keysCompare[string(s[right])]+=1
		}
		for mapCompare(keys,keysCompare)&&left<=right {
			//if (r - l + 1 < len) {
			//				len = r - l + 1
			//				ansL, ansR = l, l + len
			//			}
			if (right-left) < (ansright-ansleft){
				ansright=right+1
				ansleft=left
			}
			//if _, ok := ori[s[l]]; ok {
			//	cnt[s[l]] -= 1
			//}
			//l++
			if _,ok:=keys[string(s[left])];ok{
				keysCompare[string(s[left])]-=1
			}
			left++
		}
	}
	if ansleft==-1{
		return ""
	}
	return s[ansleft:ansright]
}

func mapCompare(m1 ,m2 map[string]int)bool{
	for k,v:=range m1{
		if v>m2[k]{
			return false
		}
	}
	return true
}