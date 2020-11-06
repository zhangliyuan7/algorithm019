package practice

//暴力查找
func TwoSum0(nums []int, target int) []int {
  for i:=0;i<len(nums);i++{
  	if nums[i]>target{
  		continue
	}
  	for j:=i+1;j<len(nums);j++{
  		if target-nums[i]==nums[j]{
  			return []int{i,j}
		}
	}
  }
  return nil
}
// map 想多的版本
func TwoSum1(nums []int, target int) []int {
	var numsKey = make(map[int][]int)
	for i,value:=range nums{
		if v,ok:=numsKey[value];!ok{
			numsKey[value]=[]int{i}
		}else{
			numsKey[value]=append(v,i)
		}
	}
	for k,v:=range numsKey{
		if v2,ok:=numsKey[target-k];ok{
			if len(v2)==1{
				if v[0]!=v2[0] {
					return []int{v[0], v2[0]}
				}
			}else{
				return v2
			}
		}
	}
	return nil
}

//题解 优秀的hash
func TwoSum2(nums []int, target int) []int {
	hashM:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if v,ok:=hashM[target-nums[i]];ok{
			return []int{v,i}
		}
		hashM[nums[i]]=i
	}
	return nil
}
