package Repeat_01

// 567
// s1 的全排列之一 ，是s2的子串（连续）
func checkInclusionSlidingWindow(s1 string, s2 string) bool {
	s1window:=[26]int{}
	s2window:=[26]int{}
	for i,c:=range s1{
		s1window[c-'a']+=1
		s2window[s2[i]-'a']+=1
	}
	if s1window==s2window{return true}
	n:= len(s1)
	for i:=n;i<len(s2);i++{
		s2window[s2[i]-'a']+=1
		s2window[s2[i-n]-'a']-=1
		if s1window==s2window{
			return true
		}
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	ls2:=len(s2)
	ls1:=len(s1)
	m:=[26]int{}
	for _,c:=range s1{
		m[c-'a']+=1
	}
	zero:=[26]int{}
	l:=0
	for l<ls2-ls1+1{
		if m[s2[l]-'a']!=0{
			tmp:=m
			for r:=l;r<ls2;r++{
				index:=s2[r]-'a'
				if tmp[index]==0{
					l++
					break
				}
				tmp[index]-=1
				if tmp==zero{
					return true
				}
			}
		}else{
			l++
		}
	}
	return false
}

