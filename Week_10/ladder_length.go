package Week_10

func LadderLength(beginWord string, endWord string, wordList []string) int {
	wdmap:=make(map[string]bool)
	for i:=range wordList{
		wdmap[wordList[i]]=true
	}
	if _,ok:=wdmap[endWord];!ok{
		return 0
	}
	startq:=make(map[string]bool)
	endq:=make(map[string]bool)
	startq[beginWord]=true
	endq[endWord]=true
	count:=1
	for len(startq)!=0{
		count++
		tmpq:=make(map[string]bool)
		for word:=range startq{
			for i:=0;i<len(word);i++{
				for j:='a';j<='z';j++{
					w:=word[:i]+string(j)+word[i+1:]
					// continue have to watch
					if w==word||startq[w]==true{
						continue
					}
					//did not need visited
					if _,ok:=endq[w];ok{
						return count
					}
					if _,ok:=wdmap[w];ok{
						tmpq[w]=true
					}
				}
			}
			delete(wdmap,word)
		}
		startq=tmpq
		if len(startq)>len(endq){
			startq,endq=endq,startq
		}
	}
	return 0
}