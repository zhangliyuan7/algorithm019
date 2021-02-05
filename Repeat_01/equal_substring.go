package Repeat_01

// 1208
func equalSubstring(s string, t string, maxCost int) int {
	longest:=0
	difs:=[]int{}
	for i:=range s{
		difs=append(difs,absDis(int(s[i])-int(t[i])))
	}
	cost:=0
	l:=0
	for r:=0;r<len(difs);r++{
		cost+=difs[r]
		for cost>maxCost{
			cost-=difs[l]
			l++
		}
		// 下面的最长获取，如果放在循环体里 有可能执行不到
		// r-l +1 (从0开始的索引，所以要加一)
		longest=maxlongest(longest,r-l+1)
	}
	return longest
}

// func equalSubstring(s string, t string, maxCost int) (maxLen int) {
//     n := len(s)
//     diff := make([]int, n)
//     for i, ch := range s {
//         diff[i] = abs(int(ch) - int(t[i]))
//     }
//     sum, start := 0, 0
//     for end, d := range diff {
//         sum += d
//         for sum > maxCost {
//             sum -= diff[start]
//             start++
//         }
//         maxLen = max(maxLen, end-start+1)
//     }
//     return
// }

func absDis(a int )int{
	if a>0{
		return a
	}
	return -a
}

func maxlongest(a,b int )int{
	if a>b{
		return a
	}
	return b
}
