package Week_08

import "fmt"

// 选择
func SelectSort(nums []int) {
	for index := 0; index < len(nums)-1; index++ {
		for i := index; i < len(nums); i++ {
			if nums[i] < nums[index] {
				nums[i], nums[index] = nums[index], nums[i]
			}
		}
	}
	fmt.Println(nums)
}

// 插入排序  不断插入到合适位置
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	fmt.Println(nums)
}

// 冒泡
func MaoPaoSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}

// 快速排序
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := 0
	r := len(nums) - 1
	mid := nums[l]
	move := 1
	for l < r {
		if nums[move] > mid {
			nums[move], nums[r] = nums[r], nums[move]
			r--
			continue
		}
		nums[l], nums[move] = nums[move], nums[l]
		l++
		move++
	}
	QuickSort(nums[0:l])
	QuickSort(nums[l+1:])
	fmt.Println(nums)
}

// 归并
func RecursionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return nums
	}
	half := len(nums) >> 1
	return merge(RecursionSort(nums[:half]), RecursionSort(nums[half:]))
}

func merge(a []int, b []int) []int {
	var r []int
	indexa := 0
	indexb := 0
	for indexa < len(a) && indexb < len(b) {
		if a[indexa] > b[indexb] {
			r = append(r, b[indexb])
			indexb++
		} else {
			r = append(r, a[indexa])
			indexa++
		}
	}
	if len(a[indexa:]) > 0 {
		r = append(r, a[indexa:]...)
	}
	if len(b[indexb:]) > 0 {
		r = append(r, b[indexb:]...)
	}
	return r
}

// 堆排序
func HeapSort(nums []int) {
	var h Heaps
	for i:=0;i<len(nums);i++{
		h.push(nums[i])
	}
	for i:=0;i<len(nums);i++{
		nums[i]=h.pop()
	}
	fmt.Println(nums)
}

type Heaps struct {
	store []int
}

func (h *Heaps) heapifyup() {
	inq := len(h.store) - 1
	father := (inq-1) /2
	for h.store[inq] < h.store[father] {
		h.store[inq], h.store[father] = h.store[father], h.store[inq]
		inq = father
		father = (inq-1)/2
	}
}

func (h *Heaps) heapifydown() {
	fa:=0
	var smaller int
	for fa*2+1<len(h.store)&&fa*2+2<len(h.store){
		smaller=fa*2+2
		if h.store[fa*2+1]<h.store[fa*2+2]{
			smaller = fa*2+1
		}
		h.store[fa],h.store[smaller]=h.store[smaller],h.store[fa]
		fa=smaller
	}
	if fa*2+1<len(h.store){
		h.store[fa],h.store[fa*2+1]=h.store[fa*2+1],h.store[fa]
	}
}

func (h *Heaps) pop() int {
	small:=h.store[0]
	h.store[0],h.store[len(h.store)-1]=h.store[len(h.store)-1],h.store[0]
	h.store=h.store[:len(h.store)-1]
	h.heapifydown()
	return small
}

func (h *Heaps) push(n int) {
	h.store = append(h.store, n)
	h.heapifyup()
}
