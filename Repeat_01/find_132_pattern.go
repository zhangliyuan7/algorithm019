package Repeat_01

import "math"

//456
func find132pattern(nums []int) bool {
	ln:=len(nums)
	leftMin:=[]int{}
	// 最小值维护，维护每一个长度的当前最小值
	for i:=0;i<ln;i++{
		leftMin=append(leftMin,math.MaxInt32)
	}
	for i:=1;i<ln;i++{
		leftMin[i]=min(leftMin[i-1],nums[i-1])
	}
	//栈内存的是可能的k值
	stack:=[]int{}
	for j:=ln-1;j>0;j--{
		numk:=math.MinInt32
		// 假如存在numj 大于 当前numk  因为从后往前遍历，栈内找到当前小于nums[j]的最大值 三者作比较即可
		for len(stack)>0&&stack[len(stack)-1]<nums[j]{
			// 保存一下当前的最大k
			numk=stack[len(stack)-1]
			// 不断下探寻求当前j时 最大的k
			stack=stack[:len(stack)-1]
		}
		// 因为上边的循环保证了numj>numk 现在如果当前numk > 0-j之间的最小值 即符合132条件
		if numk > leftMin[j]{
			return true
		}
		// j继续遍历 往前走，即可把j入到当前k队列 栈顶即为最小也就是当前的numj
		stack=append(stack,nums[j])
	}
	return false
}
