package Week_10

//
type tries struct{
	next [26]*tries
	end bool
	//word string
}

func (t *tries)searchPre(pre string )bool{
	for _,c:=range pre{
		if t.next[c-'a']!=nil{
			t=t.next[c-'a']
		}else{
			return false
		}
	}
	return true
}

func (t *tries)searchWord(w string )bool{
	for _,c:=range w{
		if t.next[c-'a']!=nil{
			t=t.next[c-'a']
		}else{
			return false
		}
	}
	return t.end
}

func (t *tries)insertWord(w string ){
	for _,c:=range w{
		if t.next[c-'a']!=nil{
			t=t.next[c-'a']
		}else{
			t.next[c-'a']=new(tries)
			t=t.next[c-'a']
		}
	}
	t.end=true
}
