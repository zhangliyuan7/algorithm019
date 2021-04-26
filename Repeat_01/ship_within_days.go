package Repeat_01

// 1011
func shipWithinDays(weights []int, D int)(ans int ){
	left:=0
	right:=0
	for i:=range weights{
		right+=weights[i]
	}
	ans=right
	for left<=right{
		mid:=left+(right-left)/2
		if CanTransfer(mid,weights,D){
			ans=min(ans,mid)
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return ans
}

func CanTransfer(n int ,weights []int ,D int )bool{
	left:=0
	for i:=0;i<len(weights);i++{
		if left+weights[i]<=n{
			left+=weights[i]
			continue
		}
		i--
		left=0
		D--
		if D<1{
			return false
		}
	}
	return true
}
