package Week_08

type LRUCache struct {
	cap  int
	length int
	hashMap map[int]*dlinklist
	head *dlinklist
	tail *dlinklist
}


func Constructor(capacity int) LRUCache {
	head:=dlinklist{pre:nil,next:nil}
	tail:=dlinklist{pre:nil,next:nil}
	head.next=&tail
	tail.pre=&head
	return LRUCache{
		cap:capacity,
		length:0,
		hashMap:make(map[int]*dlinklist),
		head:&head,
		tail:&tail,
	}
}

type dlinklist struct{
	pre *dlinklist
	next  *dlinklist
	key int
	value int
}

func (this *LRUCache) Get(key int) int {
	if v,ok:=this.hashMap[key];ok{
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if v,ok:=this.hashMap[key];ok{
		v.value=value
		this.moveToHead(v)
		return
	}
	if this.cap==this.length{
		this.popTheTail()
	}
	node:=dlinklist{pre:this.head,next:this.head.next,key:key,value:value}
	this.head.next=&node
	this.hashMap[key]=&node
	this.length++
}
func (this *LRUCache)moveToHead(d *dlinklist){
	d.pre.next=d.next
	d.next=this.head.next
	this.head.next=d
}

func (this *LRUCache)popTheTail(){
	node:=this.tail.pre
	node.pre.next=node.next
	node.next.pre=node.pre
	delete(this.hashMap,node.key)
	this.length--
}

