package Week_06

import (
	"math"
	"regexp"
)
//图方式
func ladderLengthgraph(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}


//双向广度优先
func ladderLengthd(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	distBegin := make([]int, len(wordId))
	for i := range distBegin {
		distBegin[i] = inf
	}
	distBegin[beginId] = 0
	queueBegin := []int{beginId}

	distEnd := make([]int, len(wordId))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0
	queueEnd := []int{endId}

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		q := queueBegin
		queueBegin = nil
		for _, v := range q {
			if distEnd[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distBegin[w] == inf {
					distBegin[w] = distBegin[v] + 1
					queueBegin = append(queueBegin, w)
				}
			}
		}

		q = queueEnd
		queueEnd = nil
		for _, v := range q {
			if distBegin[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distEnd[w] == inf {
					distEnd[w] = distEnd[v] + 1
					queueEnd = append(queueEnd, w)
				}
			}
		}
	}
	return 0
}

//bfs no.127  //20s bfs层序遍历 暴力求层数 too bad
// 使用正则，正则太慢了
func LadderLength(beginWord string, endWord string, wordList []string) int {
	var wordMap = make(map[string]int)
	for i := range wordList {
		wordMap[wordList[i]] = +1
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	var next_visit = make(map[string]int)
	next_visit[beginWord] = 1
	//visited:=make(map[string]int)
	var count int = 1
	for len(next_visit) != 0 {
		count += 1
		will_visited := next_visit
		next_visit = make(map[string]int)
		//fmt.Println(will_visited)
		re_string := []string{}
		for k, _ := range will_visited {
			delete(wordMap, k)
			for i, _ := range k {
				re_string = append(re_string, k[0:i]+"."+k[i+1:])
			}
		}
		for _, re := range re_string {
			for k, _ := range wordMap {
				if ok, _ := regexp.Match(re, []byte(k)); ok {
					if k == endWord {
						return count
					}
					if _, ok := next_visit[k]; !ok {
						next_visit[k] = 1
					}
				}
			}
		}
	}
	return 0
}

// 168ms 不用正则了，循环比对
func LadderLength2(beginWord string, endWord string, wordList []string) int {
	var wordMap = make(map[string]int)
	for i := range wordList {
		wordMap[wordList[i]] = +1
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	var next_visit = make(map[string]int)
	next_visit[beginWord] = 1
	delete(wordMap, beginWord)
	var count int = 1
	for len(next_visit) != 0 {
		count += 1
		will_visited := next_visit
		next_visit = make(map[string]int)
		for word := range will_visited {
			for i := 0; i < len(word); i++ {
				for r := 'a'; r <= 'z'; r++ {
					w := word[0:i] + string(r) + word[i+1:]
					if w==word{
						continue
					}
					if _, ok := wordMap[w]; ok {
						if w == endWord {
							return count
						}
						if _, ok := next_visit[w]; !ok {
							delete(wordMap, w)
							next_visit[w] = 1
						}
					}
				}
			}
		}
	}
	return 0
}
