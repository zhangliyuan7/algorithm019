package Week_07

//127 2 direct bfs
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		wordSet[wordList[i]] = true
	}
	if _,ok:=wordSet[endWord];!ok{
		return 0
	}
	startQ := make(map[string]bool)
	endQ := make(map[string]bool)
	startQ[beginWord]=true
	endQ[endWord]=true
	count := 1
	for len(startQ) != 0{
		count++
		tmpQ:=make(map[string]bool)
		for word := range startQ {
			for j := 'a'; j < 'z'+1; j++ {
				for i, _ := range word {
					new_word := word[:i] + string(j) + word[i+1:]
					if new_word==word|| startQ[new_word]==true {
						continue
					}
					if _, ok := endQ[new_word]; ok {
						return count
					}
					if _, ok := wordSet[new_word]; ok {
						tmpQ[new_word]=true
					}
				}
			}
			startQ=tmpQ
			delete(wordSet,word)
		}
		if len(startQ) > len(endQ) {
			startQ, endQ = endQ, startQ
		}
	}
	return 0
}
