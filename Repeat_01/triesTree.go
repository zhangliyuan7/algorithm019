package Repeat_01


type TriesTree struct{
	next [26]*TriesTree
	isEnd bool
	word string
}

func (t *TriesTree)SearchPrefix(prefix string )bool{
	for _,c:=range prefix{
		if t.next[c-'a']==nil{
			return false
		}
		t=t.next[c-'a']
	}
	return true
}

func (t *TriesTree)SearchWord(word string )(bool,string){
	for _,c:=range word{
		if t.next[c-'a']==nil{
			return false,""
		}
		t=t.next[c-'a']
	}
	return t.isEnd,t.word
}

func (t *TriesTree)InsertWord(word string ){
	for _,c:=range word{
		if t.next[c-'a']!=nil{
			t=t.next[c-'a']
			continue
		}
		tn:=new(TriesTree)
		tn.next=[26]*TriesTree{}
		t.next[c-'a']=tn
		t=tn
	}
	t.word=word
	t.isEnd=true
}