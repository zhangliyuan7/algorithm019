package Week_02

//双端队列 写法最简单
func maxSlidingWindowC(nums []int, k int) []int {
	ret := make([]int,0)
	stack := make([]int,0)
	for i:=0; i<len(nums); i++ {
		if i>=k && nums[i-k] == stack[0] {
			stack = stack[1:]
		}
		for len(stack) != 0 && stack[len(stack)-1]<nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,nums[i])
		if i >= k-1 {
			ret = append(ret,stack[0])
		}
	}
	return ret

}

// "单调队列" 结构用双端队列
func maxSlidingWindow(nums []int, k int) []int {
	q:=NewQ()
	r:=[]int{}
	for i:=0;i<len(nums);i++{
		if i<k-1{
			q.Push(nums[i])
		}else{
			q.Push(nums[i])
			r=append(r, q.Front())
			q.Pop(nums[i-k+1])
		}
	}
	return r
}

type Q struct{
	deque []int
}

func NewQ()*Q{
	return &Q{
		deque: []int{},
	}
}

func (q *Q)Pop(n int){
	if !q.empty()&&n==q.deque[0]{
		q.deque=q.deque[1:]
	}
}

func (q *Q)Push(n int ){
	if q.empty(){
		q.deque=append(q.deque,n)
		return
	}
	for !q.empty()&&q.deque[len(q.deque)-1]<n{
		q.deque=q.deque[:len(q.deque)-1]
	}
	q.deque=append(q.deque,n)
}

func (q *Q)Front()int{
		return q.deque[0]
}
func (q *Q)empty()bool{
	if len(q.deque)==0{
		return true
	}
	return false
}

// copy
type MonotonicQueue struct {
	Deque []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{}
}

func (m *MonotonicQueue) Push(num int) {
	for len(m.Deque) != 0 && m.Deque[len(m.Deque)-1] < num {
		m.Deque = m.Deque[:len(m.Deque)-1]
	}
	m.Deque = append(m.Deque, num)
}

func (m *MonotonicQueue) Pop(num int) {
	if len(m.Deque) != 0 && m.Deque[0] == num {
		m.Deque = m.Deque[1:]
	}
}

func (m *MonotonicQueue) Max() int {
	return m.Deque[0]
}

func maxSlidingWindowA(nums []int, k int) []int {
	var window MonotonicQueue
	var result []int
	for key, value := range nums {
		if key < k-1 { //先填满窗⼝的前 k - 1
			window.Push(value)
		} else {  // 窗⼝向前滑动
			window.Push(value)
			result = append(result, window.Max())
			window.Pop(nums[key-k+1])
		}
	}
	return result
}

//double pointer search max 类似暴力求解，优化了一部分
func maxSlidingWindowB(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums){
		return nil
	}
	var maxNums []int
	//max用于记录窗口中的最大值的索引
	max := -1
	for i:=0;i<=len(nums)-k;i++{
		l := i
		r := i+k-1
		//如果出去的值是最大值(的索引)，那么新的窗口需要重新查找最大值并设置
		if max == -1 || max == l-1 {
			max = getMax(nums,l,r)
		}else{
			//如果出去的值不是最大值(的索引)，则判断新进来的值是否大于当前的最大值
			//不大于则最大值(的索引)不变
			//大于则将进来的值(的索引)作为最大值(的索引)
			if nums[r] > nums[max] {
				max = r
			}
		}
		maxNums = append(maxNums,nums[max])
	}
	return maxNums
}

//双指针查找最大值的索引
func getMax(nums []int,left,right int) int {
	for left<right{
		if nums[left]>nums[right]{
			right--
		}else{
			left++
		}
	}
	return left
}

func maxSlidingWindowX(nums []int, k int) []int {
	ret := make([]int,0)
	stack := make([]int,0)
	for i:=0; i<len(nums); i++ {
		if i>=k && nums[i-k] == stack[0] {
			stack = stack[1:]
		}
		for len(stack) != 0 && stack[len(stack)-1]<nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,nums[i])
		if i >= k-1 {
			ret = append(ret,stack[0])
		}
	}
	return ret

}

//注意是第二个开始就
func MaxSlidingWindowA(nums []int, k int)[]int{
   var s []int
   var  r []int
   for i:=0;i<len(nums);i++{
   		if i<k-1{
			s=push(s,nums[i])
		}
		if i>=k-1{
			s=push(s,nums[i])
			r=append(r,front(s))
			s=pop(s,nums[i-k+1])
		}
   }
   return r
}

func push(a []int ,n int )[]int{
	for len(a)!=0&&a[len(a)-1]<n{
		a=a[:len(a)-1]
	}
	a=append(a,n)
	return a
}
func pop(a []int,n int )[]int{
	if a[0]==n{
		a=a[1:]
	}
	return a
}

func front(a []int )int {
	return a[0]
}














func MaxSlidingWindowD(nums []int, k int) []int {
	var maxlist []int
	s:=[]int{}
	for i:=0;i<len(nums);i++{
		if i<k{
			s=push2(s,nums[i])
		}else{
			maxlist=append(maxlist,s[0])
			if s[0]==nums[i-k]{
				s=s[1:]
			}
			s=push2(s,nums[i])
		}
	}
	maxlist=append(maxlist,s[0])
	return maxlist
}

func push2(l []int, n int )[]int{
	i:=len(l)-1
	for ;i>=0;i--{
		if l[i]>=n{
			break
		}
	}
	l=l[0:i+1]
	l=append(l,n)
	return l
}