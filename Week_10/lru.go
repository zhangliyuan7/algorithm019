package Week_10


//lru
type lru struct{
	searchkey map[int]*linkedlist
	//searchkey map[key]node
	cap int
	length int
	head *linkedlist
	tail *linkedlist
}

type linkedlist struct{
	pre *linkedlist
	next *linkedlist
	k int
	v int
}

func (l *lru)moveHead(n *linkedlist){
	//n的前后互相连接
	pre:=n.pre
	next:=n.next
	pre.next=next
	next.pre=pre
	//原head.next 与n互相连接
	n.next=l.head.next
	l.head.next.pre=n
	// head 与n互相连接
	l.head.next=n
	n.pre=l.head
}

func (l *lru)moveOut(){
	//last 是待删除
	last:=l.tail.pre
	last.pre.next=l.tail
	l.tail.pre=last.pre
	l.length-=1
}

func (l *lru)Get(k int )int {
	if v,ok:=l.searchkey[k];ok{
		l.moveHead(v)
		return v.v
	}
	return -1
}

func (l *lru)Put(key,value int ){
	if v,ok:=l.searchkey[key];ok{
		v.v=value
		l.moveHead(v)
		return
	}
	n:=new(linkedlist)
	n.k=key
	n.v=value
	//连接之前的head.next 与n
	mid:=l.head.next
	mid.pre=n
	n.next=mid
	//连接head与n
	n.pre=l.head
	l.head.next=n
	//
	l.length+=1
	//超容量淘汰tail.pre
	if l.length>l.cap{
		l.moveOut()
	}
}