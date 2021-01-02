package Week_09

//239
func MaxSlidingWindow(nums []int, k int) []int{
	q:=Queue{}
	r:=[]int{}
	for i:=0;i<len(nums);i++{
		if i<k-1{
			q.push(nums[i])
		}else{
			q.push(nums[i])
			r=append(r,q.top())
			q.pop(nums[i-k+1])
		}
	}
	return r
}

type Queue struct{
	q []int
}

func (q *Queue)push(i int ){
	if len(q.q)==0{
		q.q=append(q.q,i )
		return
	}
	for len(q.q)>0&&i>q.q[len(q.q)-1]{
		q.q=q.q[:len(q.q)-1]
	}
	q.q=append(q.q,i)
}
func (q *Queue)pop(n int ){
	if len(q.q)!=0&&q.q[0]==n{
		q.q=q.q[1:]
	}
}
func (q *Queue)top()int{
	return q.q[0]
}