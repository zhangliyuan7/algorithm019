package Repeat_01

import "container/list"

type MyHashMap struct {
	buckets [][]*Node
}

type Node struct{
	val int
}


/** Initialize your data structure here. */
func Constructorx() MyHashMap {
	buckets:=make([][]*Node,64)
	for i:=range buckets{
		buckets[i]=make([]*Node,10000)
	}
	return MyHashMap{
		buckets:buckets,
	}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	bucket:=key%64
	index:=key/64+bucket%64
	if index>=len(this.buckets[bucket]){
		tmp:=this.buckets[bucket]
		newList:=make([]*Node,index+1)
		copy(newList,tmp)
		this.buckets[bucket]=newList
	}
	if  this.buckets[bucket][index]==nil{
		this.buckets[bucket][index]=new(Node)
	}
	this.buckets[bucket][index].val=value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	bucket:=key%64
	index:=key/64+bucket%64
	if index>=len(this.buckets[bucket]){return -1 }
	if this.buckets[bucket][index]==nil{return -1 }
	return this.buckets[bucket][index].val
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	bucket:=key%64
	index:=key/64+bucket%64
	if index<len(this.buckets[bucket]){
		this.buckets[bucket][index]=nil
	}
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */


const base = 769

type entry struct {
	key, value int
}

type MyHashMap2 struct {
	data []list.List
}

func Constructor2() MyHashMap2 {
	return MyHashMap2{make([]list.List, base)}
}

func (m *MyHashMap2) hash(key int) int {
	return key % base
}

func (m *MyHashMap2) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap2) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap2) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
