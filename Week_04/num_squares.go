package Week_04

import "math"

// better one
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

// 279
//dp[i]=min(dp[i-1]+1,dp[i-j*j]+1...) 通过 ，但是很慢，420ms+
func NumSquaresDP(n int) int {
	sq:=make(map[int]int)
	for i:=1;i<n;i++{
		if i*i<=n{
			sq[i]=i*i
		}
	}
	dp:=[]int{0,1};minx:=0
	for i:=2;i<=n;i++{
		tmp:=dp[i-1]+1
		for _,v:=range sq{
			if v<=i{
				minx=min(tmp,dp[i-v]+1)
				if minx < tmp{
					tmp=minx
				}
			}else{
				break
			}
		}
		dp=append(dp,minx)
	}
	return dp[len(dp)-1]
}

//拉格朗日 四平方数 和三平方数定理解法
func numSquaresMath(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt * sqrt == n {
		return 1
	}

	for n % 4 == 0 {
		n /= 4
	}
	if n % 8 == 7 {
		return 4
	}

	for i := 0; i * i < n; i++ {
		s := int(math.Sqrt(float64(n-i*i)))
		if s * s == n - i * i {
			return 2
		}
	}
	return 3
}


func min(a,b int )int {
	if a<b{
		return a
	}
	return b
}
// 找开方数。。。bad 合理做法是动态规划
func NumSquares(n int) int {
	sqrt:=0
	count:=0
	for {
		sqrt=SearchSqrt(n)
		count+=1
		if sqrt*sqrt==n{
			break
		}
		if sqrt*sqrt==n-(sqrt*sqrt){
			count+=1
			break
		}
		n=n-(sqrt*sqrt)
	}
	return count
}
func SearchSqrt(n int )int {
	l,r:=0,n
	for l<=r {
		mid:=l+(r-l)/2
		sq:=mid*mid
		if sq==n{
			return mid
		}
		if sq<n{
			l=mid+1
		}
		if sq>n{
			r=mid-1
		}
	}
	return l-1
}
