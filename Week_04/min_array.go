package Week_04

// erfen good
func MinArray1(numbers []int) int {
	left,right:=0,len(numbers)-1
	for left<right {
		mid:=left+(right-left)/2
		if numbers[mid]>numbers[right]{
			left=mid+1
		}else{
			right--
		}
	}
	return numbers[left]
}

//erfen
func MinArray(numbers []int) int {
	left,right := 0,len(numbers)-1
	if len(numbers)==1{
		return numbers[0]
	}
	if numbers[len(numbers)-1]>numbers[0]{
		return numbers[0]
	}
	for left<right{
		mid:=left+(right-left)/2
		if numbers[mid]<numbers[right]{
			right=mid
			continue
		}
		if numbers[mid]>numbers[right]{
			left=mid+1
			continue
		}
		if numbers[mid]==numbers[right]{
			right--
		}
	}
	return numbers[left]
}
