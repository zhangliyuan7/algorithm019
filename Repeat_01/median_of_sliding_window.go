package Repeat_01

// 480

//type DHeap struct {
//	large     *smallTopHeap
//	small     *smallTopHeap
//	delayDrop map[int]int
//}
//
//func NewDHeap() *DHeap {
//	dh := DHeap{}
//	dh.large = NewSmallTopHeap()
//	dh.small = NewSmallTopHeap()
//	dh.delayDrop = make(map[int]int)
//	return &dh
//}
//
//func (dh *DHeap) Insert(x int) {
//	fmt.Println("insert : ", x)
//	if dh.small.size == 0 {
//		dh.small.Insert(-x)
//		return
//	}
//	if -dh.small.Top() >= x {
//		dh.small.Insert(-x)
//	} else {
//		dh.large.Insert(x)
//	}
//	dh.MakeBalance()
//}
//
//func (dh *DHeap) cleanDelay() {
//	for k, v := range dh.delayDrop {
//		if v > 0 {
//			if dh.small.Top() == -k {
//				dh.small.Pop()
//				dh.small.size++
//			} else if dh.large.Top() == k {
//				dh.large.Pop()
//				dh.large.size++
//			}
//			dh.delayDrop[k]--
//		} else {
//			delete(dh.delayDrop, k)
//		}
//	}
//}
//
//func (dh *DHeap) MakeBalance() {
//	dh.cleanDelay()
//	if dh.large.size > dh.small.size {
//		dh.small.Insert(-dh.large.Pop())
//	} else if dh.large.size+1 < dh.small.size {
//		dh.large.Insert(-dh.small.Pop())
//	}
//	dh.cleanDelay()
//}
//
//func (dh *DHeap) Pop(x int) {
//	if dh.large.Top() != x && dh.small.Top() != -x {
//		dh.delayDrop[x] += 1
//		if dh.large.Top() < x {
//			dh.large.size -= 1
//		} else {
//			dh.small.size -= 1
//		}
//	} else {
//		if dh.small.Top() == -x {
//			dh.small.Pop()
//
//		} else if dh.large.Top() == x {
//			dh.large.Pop()
//		}
//	}
//	dh.MakeBalance()
//}
//
//func (dh *DHeap) GetMedian(k int) float64 {
//	if k&1 == 1 {
//		return -float64(dh.small.Top())
//	}
//	return float64(-dh.small.Top()+dh.large.Top()) / 2
//}
//
//type smallTopHeap struct {
//	size  int
//	store []int
//}
//
//func NewSmallTopHeap() *smallTopHeap {
//	return &smallTopHeap{0, []int{}}
//}
//
//func (sth *smallTopHeap) Insert(x int) {
//	sth.store = append(sth.store, x)
//	sth.size++
//	sth.heapifyUp()
//}
//
//func (sth *smallTopHeap) Pop() int {
//	r := sth.store[0]
//	sth.store[0] = sth.store[len(sth.store)-1]
//	sth.store = sth.store[:len(sth.store)-1]
//	sth.size--
//	sth.heapifyDown()
//	return r
//}
//
//func (sth *smallTopHeap) Top() int {
//	return sth.store[0]
//}
//
//func (sth *smallTopHeap) heapifyUp() {
//	ch := len(sth.store) - 1
//	fa := (ch - 1) / 2
//	for sth.store[fa] > sth.store[ch] {
//		sth.store[fa], sth.store[ch] = sth.store[ch], sth.store[fa]
//		ch = fa
//		fa = (ch - 1) / 2
//	}
//}
//
//func (sth *smallTopHeap) heapifyDown() {
//	fa := 0
//	ch1 := 1
//	ch2 := 2
//	if len(sth.store) == 2 && sth.store[fa] > sth.store[ch1] {
//		sth.store[fa], sth.store[ch1] = sth.store[ch1], sth.store[fa]
//		return
//	}
//	for ch1 < len(sth.store) && ch2 < len(sth.store) {
//		if sth.store[ch1] > sth.store[ch2] {
//			sth.store[fa], sth.store[ch2] = sth.store[ch2], sth.store[fa]
//			fa = ch2
//		} else {
//			sth.store[fa], sth.store[ch1] = sth.store[ch1], sth.store[fa]
//			fa = ch1
//		}
//		ch1 = fa*2 + 1
//		ch2 = fa*2 + 2
//	}
//	//fmt.Println("heapifydown end ",sth.store)
//
//}
//
//func medianSlidingWindowReDo(nums []int, k int) []float64 {
//	r := []float64{}
//	dh := NewDHeap()
//	for i := 0; i < k; i++ {
//		dh.Insert(nums[i])
//	}
//
//	r = append(r, dh.GetMedian(k))
//	for i := k; i < len(nums); i++ {
//		//fmt.Println("large - ",dh.large.store)
//		//fmt.Println("small - ",dh.small.store)
//		dh.Insert(nums[i])
//		dh.Pop(nums[i-k])
//		//fmt.Println("large - ",dh.large.store)
//		//fmt.Println("small - ",dh.small.store)
//		//fmt.Println("\n-------------")
//		r = append(r, dh.GetMedian(k))
//	}
//	return r
//}
