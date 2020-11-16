package w2

import "sort"

//暴力 28ms
func groupAnagrams(strs []string) [][]string {
	var r = [][]string{}
	h:=make(map[[26]int][]string)
	for _,v:=range strs{
		var tmpArr = [26]int{}
		for x:=0;x<len(v);x++{
			tmpArr[v[x]-'a']++
		}
		h[tmpArr]=append(h[tmpArr],v)
	}
	for _,v:=range h{
		r=append(r,v)
	}
	return r
}

//优化一点点 24ms
func groupAnagramsA(strs []string) [][]string {
	var r = [][]string{}
	h:=make(map[[26]int]int )
	for _,v:=range strs{
		var tmpArr = [26]int{}
		for x:=0;x<len(v);x++{
			tmpArr[v[x]-'a']++
		}
		if index,ok:=h[tmpArr];ok{
			r[index]=append(r[index],v)
		}else{
			r=append(r,[]string{v})
			h[tmpArr]=len(r)-1
		}
	}
	return r
}

//leetcode resolve copy  20ms  use sort
type bytes []byte
// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i,j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func groupAnagramsB(strs []string) [][]string {
	ret := [][]string{}
	m := make(map[string]int)
	for _, str := range strs {
		kByte := bytes(str)
		sort.Sort(kByte) // 将字符串排序
		k := string(kByte)
		if idx, ok := m[k]; !ok {
			m[k] = len(ret) // 记录拍完序的字符串以及字符串在ret的位置
			ret = append(ret, []string{str})
		} else { // 已经出现过，将其放入出现在ret中，在ret的位置为idx
			ret[idx] = append(ret[idx], str)
		}
	}
	return ret
}

//异位词 两种比较 一种是列表hash存词， 一种是重新排序词对比
