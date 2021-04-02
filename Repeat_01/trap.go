package Repeat_01

// 单调栈
func trap(height []int) (area int ){
	s:=[]int{}
	for i,k:=range height{
		if len(s)==0{
			s=append(s,i)
		}
		for len(s)>0&&height[s[len(s)-1]]<k{
			topIndex:=s[len(s)-1]
			s=s[:len(s)-1]
			if len(s)==0{
				break
			}
			leftIndex:=s[len(s)-1]
			width:=i-leftIndex-1
			height:=minTrap(height[leftIndex],k)-height[topIndex]
			area+=width*height
		}
		s=append(s,i)
	}
	return
}

//dp min (rightmax ,leftmax ) - h
func trapII(height []int) (area int ) {
	ln:=len(height)
	leftMax:=make([]int,ln)
	leftMax[0] = height[0]
	for i:=1;i<ln;i++{
		leftMax[i]=maxTrap(leftMax[i-1],height[i])
	}
	rightMax:=make([]int,ln)
	rightMax[ln-1]=height[ln-1]
	for i:=ln-2;i>=0;i--{
		rightMax[i]=maxTrap(rightMax[i+1],height[i])
	}
	for i:=0;i<ln;i++{
		area+=minTrap(leftMax[i],rightMax[i])-height[i]
	}
	return
}

func maxTrap(a,b int )int {
	if a<b{
		return b
	}
	return a
}

func minTrap(a,b int )int {
	if a>b{
		return b
	}
	return a
}

// 双指针 最快 0ms
func trapIII(height []int) (area int ) {
	ln:=len(height)
	if ln<=1{
		return 0
	}
	left:=0
	right:=ln-1
	leftMax:=0
	rightMax:=0
	for left<right{
		leftMax=maxTrap(leftMax,height[left])
		rightMax=maxTrap(rightMax,height[right])
		if height[left]<height[right]{
			area+=leftMax-height[left]
			left++
		}else {
			area+=rightMax-height[right]
			right--
		}
	}
	return
}
