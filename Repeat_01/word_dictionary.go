package Repeat_01

// 211
type WordDictionary struct {
	next [26]*WordDictionary
	isEnd bool
	word string
}


/** Initialize your data structure here. */
func ConstructorWD() WordDictionary {
	wd:=WordDictionary{}
	wd.next =[26]*WordDictionary{}
	return wd
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	for _,c:=range word{
		if this.next[c-'a']!=nil{
			this=this.next[c-'a']
			continue
		}
		wd:=new(WordDictionary)
		wd.next=[26]*WordDictionary{}
		this.next[c-'a']=wd
		this=wd
	}
	this.isEnd=true
	this.word=word
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for i,c:=range word{
		if c=='.'{
			for _,v:=range this.next{
				newword:=word[i+1:]
				//这里注意 当前的v是下层的v ，当前的newword是不包含当前匹配.位置的新word
				//递归调用，注意是下层对象，而非this
				if v!=nil&&v.Search(newword)==true {
					return true
				}
			}
			return false
		}
		if this.next[c-'a']==nil{
			return false
		}
		this=this.next[c-'a']
	}
	return this.isEnd
}
