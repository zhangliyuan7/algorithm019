package Repeat_01

// 978
func maxTurbulenceSize(arr []int) int {
  dp:=make([][]int,len(arr))
  for i:=range dp{
  	dp[i]=[]int{}
  }
  dp[0]=[]int{1,1}
  for i:=1;i<len(arr);i++{
  	dp[i]=[]int{1,1}
  	if arr[i]>arr[i-1]{
  		dp[i][0]=dp[i-1][1]+1
	}else if arr[i]<arr[i-1]{
		dp[i][1]=dp[i-1][0]+1
	}
  }
  ans:=0
  for i:=range dp{
  	ans=max(ans,dp[i][0])
  	ans=max(ans,dp[i][1])
  }
  return ans
}

// sliding window
func maxTurbulenceSizeSlidingWindow(arr []int) int {
	ln:=len(arr)
	ml:=1
	l,r:=0,0
	for r < ln-1{
		if l==r{
			if arr[l]==arr[l+1]{
				l++
			}
			r++
		}else{
			if arr[r]>arr[r+1] && arr[r]> arr[r-1] {
				r++
			} else if arr[r] < arr[r+1] && arr[r] <arr[r-1] {
				r++
			} else {
				l = r
			}
		}
		ml=max_maxtur(ml,r-l+1)
	}
	return ml
}
func max_maxtur(a, b int )int{
	if a>b{
		return a
	}
	return b
}