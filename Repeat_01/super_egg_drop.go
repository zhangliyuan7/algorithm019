package Repeat_01

import "math"

//887
// 用m作为最少检查次数，判断可以判断的楼层高度，
// k m 可分为两种情况
// 1 - 鸡蛋碎了 k-1 ，可尝试次数少了一次 m-1
// 2 - 鸡蛋没碎 k，但是可尝试次数少了一次 m-1
// 3 - 判断出F的楼层
func superEggDropM(K int ,N int )int{
	dp:=make([][]int,K+1)
	for i:=0;i<K+1;i++{
		dp[i]=make([]int,N+1)
	}
	for m:=1;m<=N;m++{
		for k:=1;k<=K;k++{
			dp[k][m]=dp[k-1][m-1]+dp[k][m-1]+1
			if dp[k][m]>=N{
				return m
			}
		}
	}
	return N
}

// second way has some problems with maxInt32
func superEggDropX(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= K; i++ {
		dp[i][1] = 1
	}
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			for x := 1; x <= j; x++ {
				if dp[i][j] == 0 {
					dp[i][j] = math.MaxInt32
				}
				dp[i][j] = min(dp[i][j], max(dp[i][j-x], dp[i-1][x-1])+1)
			}
		}
	}
	return dp[K][N]
}

//overflow timeout
func superEggDrop(K int, N int) int {
	m:=make(map[[2]int]int)
	var dp func(k int,n int )int
	dp =func(k int ,n int )int{
		if k==1{
			return n
		}
		if n==0{
			return 0
		}
		if v,ok:=m[[2]int{k,n}];ok{
			return v
		}
		res:=math.MaxInt32
		for i:=0;i<n+1;i++ {
			res = min(res, max(superEggDrop(k-1 ,i-1), superEggDrop(k, n-i))+1)
		}
		m[[2]int{k,n}]=res
		return res
	}

	return dp(K,N)
}

//overflow
func superEggDropLR(K int ,N int )int{
	dict:=make(map[[2]int]int )

	var dp  func(k int,n int)int

	dp= func(k int, n int) int {
		if k==1{
			return n
		}
		if n==0{
			return 0
		}
		if v,ok:=dict[[2]int{k,n}];ok{
			return v
		}
		res:=math.MaxInt32
		lo,hi:=1,n
		for lo<=hi{
			mid:=lo+(hi-lo)/2
			broken:=dp(k-1,mid-1)
			notBroken:=dp(k,N-mid)
			if broken>notBroken{
				hi=mid-1
				res=min(broken+1,res)
			}else{
				lo=mid+1
				res=min(notBroken+1,res)
			}
		}
		dict[[2]int{k,n}]=res
		return res
	}
	return dp(K,N)
}