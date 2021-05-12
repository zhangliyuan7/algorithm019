package Repeat_01

//1310  前缀和原理 + 同项异或为0
func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = xors[q[0]] ^ xors[q[1]+1]
	}
	return ans
}


func xorQueriesI(arr []int, queries [][]int) (ans []int) {
	sums:=make([]int,len(arr))
	sums[0]=arr[0]
	for i:=1;i<len(arr);i++{
		sums[i]=sums[i-1]^arr[i]
	}
	ans=make([]int,len(queries))
	for i:=range queries{
		if queries[i][0]!=0{
			ans[i]=sums[queries[i][1]]^sums[queries[i][0]-1]
		}else{
			ans[i]=sums[queries[i][1]]
		}
	}
	return
}