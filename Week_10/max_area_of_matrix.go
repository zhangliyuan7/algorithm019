package Week_10

// 84
func LargestRectangleArea(heights []int) int {
	s:=[]int{}
	n:=len(heights)
	left:=make([]int,n)
	right:=make([]int,n)
	// search left limit
	for i:=0;i<n;i++{
		for  len(s)>0&&heights[s[len(s)-1]]>=heights[i]{
			s=s[:len(s)-1]
		}
		if len(s)==0{
			left[i]=-1
		}else{
			left[i]=s[len(s)-1]
		}
		s=append(s,i)
	}
	//right limit
	s=[]int{}
	for i:=n-1;i>=0;i--{
		for len(s)>0&&heights[s[len(s)-1]]>=heights[i]{
			s=s[:len(s)-1]
		}
		if len(s)==0{
			right[i]=n
		}else{
			right[i]=s[len(s)-1]
		}
		s=append(s,i)
	}
	area:=0
	for i:=0;i<n;i++{
		area=maxArea(area,(right[i]-left[i]-1)*heights[i])
	}
	return area
}
func maxArea(a,b int )int{
	if a >b{
		return a
	}
	return b
}