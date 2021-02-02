package Repeat_01

// 424
func characterReplacement(s string, k int) int {
	counts:=[26]int{}
	maxN := 0
	l:=0
	r:=0
	for ;r<len(s);r++{
		// 记录当前字母的次数
		counts[s[r]-'A']++
		// 获取当前数量最多次出现的字母的个数
		maxN=max(maxN,counts[s[r]-'A'])
		// r-l+1 窗口长度 ，- maxN 当前最多次出现的字母，如果大于k ，那么其他字母的和就超过了容错的k个，那么窗口要缩小
		// 另当前窗口保留新元素所在位置代表的长度，所以left要缩小1 适配长度变化
		// 注意 maxN是变化的
		if r-l+1-maxN>k{
			// sliding window left
			counts[s[l]-'A']--
			l++
		}
	}
	//r==len(s),所以都无所谓是r-l 或 len(s)-l
	return r-l
}
