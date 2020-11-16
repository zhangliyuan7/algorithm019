package w2

import "sort"

//baoli
func Intersect(nums1 []int, nums2 []int) []int {
	numsR:=[]int{}
	for i:=0;i<len(nums1);i++{
		for j:=0;j<len(nums2);j++{
			if nums2[j]==nums1[i]{
				numsR=append(numsR,nums1[i])
				nums2[j],nums2[len(nums2)-1]=nums2[len(nums2)-1],nums2[j]
				nums2=nums2[:len(nums2)-1]
				break
			}
		}
	}
	return numsR
}


//double pointer
func IntersectA (nums1 []int ,nums2 []int )[]int {
	numsR:=[]int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i,j int
	for i<len(nums1)&&j<len(nums2){
		if nums1[i]==nums2[j]{
			numsR=append(numsR,nums1[i])
			i++
			j++
			continue
		}
		if nums1[i]< nums2[j]{ i++;continue }
		if nums1[i]> nums2[j]{ j++;continue }
	}
	return numsR
}

//hash map
func IntersectB(nums1 []int ,nums2[]int )[]int {
	numsR:=[]int{}
	a:=make(map[int ]int )
	for i:=0;i<len(nums1);i++{
		if _,ok:=a[nums1[i]];ok{
			a[nums1[i]]+=1
		}else{
			a[nums1[i]]=1
		}
	}
	for j:=0;j<len(nums2);j++{
		if v,ok:=a[nums2[j]];ok&&v!=0 {
			numsR = append(numsR, nums2[j])
			a[nums2[j]]-=1
		}
	}
	return numsR
}