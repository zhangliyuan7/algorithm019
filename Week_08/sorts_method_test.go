package Week_08

import (
	"testing"
)

var nums = []int{3,5,1,77,12,5,12,3,1,0}

//func TestSelectSort(t *testing.T) {
//	SelectSort(nums)
//}

//func TestInsertSort(t *testing.T){
//	InsertSort(nums)
//}

//func TestMaoPaoSort(t *testing.T) {
//	MaoPaoSort(nums)
//}

//func TestQuickSort(t *testing.T) {
//	QuickSort(nums)
//}

//func TestRecursionSort(t *testing.T) {
//	fmt.Println(RecursionSort(nums))
//}

func TestHeapSort(t *testing.T) {
	HeapSort(nums)
}