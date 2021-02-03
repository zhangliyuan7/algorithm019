package Repeat_01

// 480


type  tinyTopHeap struct{
	s []int
	size int
}

func (b *tinyTopHeap)insert(x int ){
	b.s=append(b.s,x)
	b.size++
	b.heapifyup()
}

func (b *tinyTopHeap)Top()int{
	if len(b.s)!=0{
		return b.s[0]
	}
	return 0
}

func (b *tinyTopHeap)pop()int{
	r:=b.s[0]
	b.s[0]=b.s[len(b.s)-1]
	b.s=b.s[:len(b.s)-1]
	b.size--
	b.heapifydown()
	return r
}

func (b *tinyTopHeap)heapifyup(){
	if len(b.s)==1{ return }
	last:=len(b.s)-1
	fa:=(last-1)/2
	if b.s[fa]>b.s[last]{
		b.s[fa],b.s[last]=b.s[last],b.s[fa]
		last=fa
		fa=(last+1)/2
	}
}

func (b *tinyTopHeap)heapifydown(){
     fa:=0
     ch1:=fa*2+1
     ch2:=fa*2+2
     for ch1<len(b.s)&&ch2<len(b.s){
     	if b.s[ch1]<b.s[ch2]{
            b.s[ch1],b.s[fa]=b.s[fa],b.s[ch1]
            fa=ch1
		}else{
			b.s[ch2],b.s[fa]=b.s[fa],b.s[ch2]
			fa=ch2
		}
		ch1=fa*2+1
		ch2=fa*2+2
	 }
}

func (b *tinyTopHeap) prune(delayed map[int]int,sign int) {
	for len(b.s) > 0 {
		num := b.s[0]
		if sign == -1 {
			num = -num
		}
		if d, has := delayed[num]; has {
			b.size--
			if d > 1 {
				delayed[num]--
			} else {
				delete(delayed, num)
			}
			b.pop()
		} else {
			break
		}
	}
}

func makeBalance(small *tinyTopHeap,large *tinyTopHeap,delayed map[int]int){
		if small.size>large.size+1 {
			large.insert(-small.pop())
			small.prune(delayed,-1)
		}else if small.size<large.size{
			small.insert(-large.pop())
			large.prune(delayed,1)
		}
}

func InsertToHeap(x int,small *tinyTopHeap,large *tinyTopHeap,delayed map[int]int){
		if len(small.s) == 0 || x <= -small.s[0] {
			small.insert(-x)
		} else {
			large.insert(x)
		}
		makeBalance(small,large,delayed)
	}

func erase (x int,small ,large *tinyTopHeap,delayed map[int]int ){
	delayed[x]++
	if x <= -small.s[0] {
		if x == -small.s[0] {
			small.prune(delayed,-1)
		}
	} else {
		if x == large.s[0] {
			large.prune(delayed,1)
		}
	}
	makeBalance(small,large,delayed)
}

func getMedian(k int ,small,large *tinyTopHeap) float64 {
	if k&1 > 0 {
		return float64(-small.s[0])
	}
	return float64(-small.s[0]+large.s[0]) / 2
}


func medianSlidingWindow(nums []int, k int) []float64 {
	delay:=make(map[int]int)
	small :=&tinyTopHeap{[]int{},0}
	large:= &tinyTopHeap{[]int{},0}

	for _, num := range nums[:k] {
		InsertToHeap(num,small,large,delay)
	}
	ans:=[]float64{}
	ans=append(ans,getMedian(k,small,large))
	for i:=k;i<len(nums);i++{
		InsertToHeap(nums[i],small,large,delay)
		erase(nums[i-k],small,large,delay)
		ans = append(ans, getMedian(k,small,large))
	}
	return ans
}
