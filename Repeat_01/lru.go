package Repeat_01

type lru struct {
	m map[string]*node
	head *node
	tail *node
	length int
	cap int
}

type node struct{
	pre *node
	next *node
	key string
	value int
}

func constructor(capacity int) lru {
	l:=lru{}
	l.head=new(node)
	l.tail=new(node)
	l.cap=capacity
	l.length=0
	l.m=make(map[string]*node)
	//
	l.head.next=l.tail
	l.tail.pre=l.head
	return l
}

func (l *lru )get(key string )int {
	if v,ok:=l.m[key];ok{
		l.movetohead(v)
		return v.value
	}
	return -1
}

func (l *lru)push(key string ,value int ){
	if v,ok:=l.m[key];ok{
		v.value=value
		l.movetohead(v)
		return
	}
	n:=new(node)
	n.key=key
	n.value=value
	l.m[key]=n
	l.length+=1
	if l.length>l.cap{
		l.moveout()
	}
	n.pre=l.head
	n.next=l.head.next
	l.head.next=n
	n.next.pre=n
}

func (l *lru)movetohead(n *node){
	//old location
	nnext:=n.next
	npre:=n.pre
	npre.next=nnext
	nnext.pre=npre

	// new location
	hnext:=l.head.next
	hnext.pre=n
	n.next=hnext
	n.pre=l.head
	l.head.next=n
}

func (l *lru)moveout(){
	last:=l.tail.pre
	lastpre:=last.pre
	lastpre.next=l.tail
	l.tail.pre=lastpre
	delete(l.m,last.key)
}