package practice

import "math"

func  maxArea(height []int) int {
	max:=0
	left := 0
	right :=len(height)-1
	for left<right{
		if height[left]>height[right]{
			area:=height[right]*(right-left)
			max=maxV(area,max)
			right--
		}else{
			area:=height[left]*(right-left)
			max=maxV(area,max)
			left++
		}
	}
	return max
}

func maxV(a,b int)int {
	if a>b{
		return a
	}
	return b
}

func MaxArea0(height []int) int {
	max:=0
	left := 0
	right :=len(height)-1
	for left<right{
		area:=min(height[left],height[right])*(right-left)
		if area>max{
			max=area
		}
		if height[left]>height[right]{
			right--
		}else{
			left++
		}
	}
	return max
}

func min(a,b int )int {
	if a>b{
		return b
	}
	return a
}
//暴力法1
func MaxArea1(height []int) int {
	var max int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := int(math.Min(float64(height[i]), float64(height[j])) * float64(j-i))
			if area > max {
				max = area
			}
		}
	}
	return max
}

//暴力法2
func MaxArea2(height []int) int {
	var max int
	var heightC int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				heightC = height[j]
			} else {
				heightC = height[i]
			}
			area := heightC * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

//
//收敛法
func MaxArea3(height []int) int {
	var (
		max    int
		area   int
		length = len(height)
		j      = length - 1
	)
	for i := 0; i < length; {
		if height[i] > height[j] {
			area = height[j] * (j - i)
			j--
		} else {
			area = height[i] * (j - i)
			i++
		}
		if max < area {
			max = area
		}
	}
	return max
}

//4
func MaxArea4(height []int) int {
	max := 0
	j := len(height) - 1
	for i := 0; i < len(height)-1; {
		if height[i] > height[j] {
			if v := height[j] * (j - i); v > max {
				max = v
			}
			j--
			continue
		}
		if v := height[i] * (j - i); v > max {
			max = v
		}
		i++
	}
	return max
}

//baoli
func MaxArea(ys []int)int{
	var max int
	for i:=0;i<len(ys)-1;i++{
		for j:=i;j<len(ys);j++{
			area:=min(ys[i],ys[j])*(j-i)
			max=maxV(area,max)
		}
	}
	return max
}
// shuangzhizhen
func MaxAreaD(height []int )int {
	var max int
	var j = len(height)-1
	for i:=0;i<len(height);{
		if height[i]>height[j]{
			max=maxV(height[j]*(j-i),max)
			j--
		}else{
			max=maxV(height[i]*(j-i),max)
			i++
		}
	}
	return max
}
//////
func MaxAreaE(height []int )int {
	var max int
	var last =len(height)-1
	for i:=0;i<len(height);{
		if height[i]<height[last]{
			area:=height[i]*(last-i)
			max=maxV(max,area)
			i++
		}
		if height[last]<height[i]{
			area:=height[last]*(last-i)
			max=maxV(max,area)
			last--
		}
		if i==last{
			break
		}
	}
	return max
}
