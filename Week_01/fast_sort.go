package practice

//good ,it's ez to understand ,first choose
func FastSortA(nums []int){
	if len(nums)<=1{
		return
	}
	var move = 1
	left:=0
	right:=len(nums)-1
	var mid = nums[left]
	for left<right{
		if nums[move]>mid{
			nums[move],nums[right]=nums[right],nums[move]
			right--
		}else{
			nums[move],nums[left]=nums[left],nums[move]
			left++
			move+=1
		}
	}
	//nums[left]=mid  // useless ,just delete

	FastSortA(nums[0:left])
	FastSortA(nums[left+1:])
}


// 类似。。。游标卡尺 标准在迁移
func Quick3Sort(a []int,left int, right int)  {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right ; i++ {
		if a[left] >= a[i]{
			//分割位定位++
			explodeIndex ++;
			a[i],a[explodeIndex] = a[explodeIndex],a[i]
		}
	}
	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex],a[left] // here hard to understand

	Quick3Sort(a,left,explodeIndex - 1)
	Quick3Sort(a,explodeIndex + 1,right)

}

//bad !!!
func FastSortB(nums []int){
	if len(nums)<=1{
		return
	}
	var move = 1
	left:=0
	right:=len(nums)-1
	var mid = left
	//mid用索引的话 值一直在变
	for left<right{
		if nums[move]>nums[mid]{
			nums[move],nums[right]=nums[right],nums[move]
			right--
		}else{
			nums[move],nums[left]=nums[left],nums[move]
			left++
			move+=1
		}
	}
	nums[left]=nums[mid]

	FastSortA(nums[0:left])
	FastSortA(nums[left+1:])
}