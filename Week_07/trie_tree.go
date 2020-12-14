package Week_07

//208
type TrieTree struct {
	Next [26]*TrieTree
	End bool
}

func NewTrieTree()TrieTree{
	return TrieTree{Next:[26]*TrieTree{}}
}

func (trieT *TrieTree)Insert(word string ){
	tmptt:=trieT
	for _,c:=range word{
		if tmptt.Next[c-'a']!=nil{
			tmptt=tmptt.Next[c-'a']
			continue
		}
		tmptt.Next[c-'a']=new(TrieTree)
		tmptt=tmptt.Next[c-'a']
	}
	tmptt.End=true
}

func (trieT *TrieTree)Search(word string)bool{
	//词标记 和 通路
	tmptt:=trieT
	for _,c:=range word{
		if tmptt.Next[c-'a']==nil{
			return false
		}
		if tmptt.Next[c-'a']!=nil{
			tmptt=tmptt.Next[c-'a']
		}
	}
	if !tmptt.End{
		return false
	}
	return true
}

func (trieT *TrieTree)StartWith(pre string)bool{
	//前缀序列检查
	tmptt:=trieT
	for _,c:=range pre{
		if tmptt.Next[c -'a']==nil{
			return false
		}
		if tmptt.Next[c-'a']!=nil{
			tmptt=tmptt.Next[c-'a']
		}
	}
	return true
}

type Trie struct {
	Next map[string]*Trie
	End bool
}
func Constructor() Trie {
	return Trie{Next:make( map[string]*Trie)}
}

func NewTrie() *Trie{
	t:=new(Trie)
	t.Next=make(map[string]*Trie)
	return t
}
func (t *Trie)Insert(word string ){
	for _,c:=range word{
		if _,ok:=t.Next[string(c)];!ok {
			t.Next[string(c)] = NewTrie()
		}
		t=t.Next[string(c)]
	}
	t.End=true
}

func (t *Trie)Search(word string )bool{
	for _,c:=range word{
		if _,ok:=t.Next[string(c)];!ok{
			return false
		}
		t=t.Next[string(c)]
	}
	if t.End{
		return true
	}
	return false
}
func (t *Trie)SearchStartWith(prefix string )bool{
	for _,c:=range prefix{
		if _,ok:=t.Next[string(c)];!ok{
			return false
		}
		t=t.Next[string(c)]
	}
	return true
}