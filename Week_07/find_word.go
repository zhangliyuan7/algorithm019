package Week_07

import (
	"encoding/json"
)

//212
//单词搜索2

// word传递太费空间也太费时间
// 灵活改变一下trie树结构 就可以快速加入回结果
// word传递效率低，超时
func FindWords(board [][]byte, words []string) []string {
	var r []string
	t:=newtries()
	for _,w:=range words{
		t.Insert(w)
	}
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			//查找每个值的深度优先遍历是否在trietree上，如果在，添加到结果
			dfs2(board,i,j,t,&r)
		}
	}
	return r
}
func dfs2(board [][]byte,i,j int ,t *Tries, r *[]string ){
	if i>len(board)-1||j>len(board[0])-1||i<0||j<0{
		return
	}
	if board[i][j]=='0'||t.Next[string(board[i][j])]==nil{
		return
	}
	dx:=[]int{1,0,-1,0}
	dy:=[]int{0,-1,0,1}
	tmp:=board[i][j]
	board[i][j]='0'
	t=t.Next[string(tmp)]
	if t.End!=""{
		*r=append(*r,t.End)
		t.End=""
	}
		for k:=0;k<4;k++{
			tmpi:=i+dx[k]
			tmpj:=j+dy[k]
			dfs2(board,tmpi,tmpj,t,r)
		}
	board[i][j]=tmp
}
type Tries struct{
	End string
	Next map[string]*Tries
}

func (t Tries)String()string{
	res,_:=json.Marshal(t.Next)
	return string(res)
}
func newtries()*Tries{
	t:=new(Tries)
	t.Next=make(map[string]*Tries)
	return t
}

func (t *Tries)Insert(word string ){
	for _,c:=range word{
		if _,ok:=t.Next[string(c)];!ok{
			t.Next[string(c)]=newtries()
		}
		t=t.Next[string(c)]
	}
	t.End=word
}

func (t *Tries)Search(word string) bool{
	for _,c:=range word{
		if _,ok:=t.Next[string(c)];!ok{
			return false
		}
		t=t.Next[string(c)]
	}
	if t.End!=""{
		return true
	}
	return false
}
func (t *Tries)SearchStartWith(prefix string)bool{
	for _,c:=range prefix{
		if _,ok:=t.Next[string(c)];!ok{
			return false
		}
		t=t.Next[string(c)]
	}
	return  true
}



func findWords(board [][]byte, words []string) []string {
	trie,res := buildTrieTree(words),make([]string,0)
	for i := 0; i < len(board); i++{
		for j := 0; j < len(board[0]); j++{
			dfsx(board,trie,i,j,&res)
		}
	}
	return res
}
func dfsx(board [][]byte, node *trieNode, i, j int, res *[]string){
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
		return
	}
	c := board[i][j]
	if board[i][j] == '#' || node.nodes[c - 'a'] == nil{
		return
	}
	node = node.nodes[c - 'a']
	if node.word != ""{
		*res = append(*res,node.word)
		node.word = ""
	}
	board[i][j] = '#'
	dfsx(board,node,i-1,j,res)
	dfsx(board,node,i+1,j,res)
	dfsx(board,node,i,j-1,res)
	dfsx(board,node,i,j+1,res)
	board[i][j] = c
}
type trieNode struct{
	word string
	nodes []*trieNode
}
func buildTrieTree(words []string)*trieNode{
	root := trieNode{nodes:make([]*trieNode,26)}
	for _, word := range words{
		cur := &root
		for i := 0; i < len(word); i++{
			c := word[i]
			if cur.nodes[c - 'a'] == nil{
				cur.nodes[c - 'a'] = &trieNode{nodes:make([]*trieNode,26)}
			}
			cur = cur.nodes[c - 'a']
		}
		cur.word = word
	}
	return &root
}