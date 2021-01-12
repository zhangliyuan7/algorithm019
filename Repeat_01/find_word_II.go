package Repeat_01

// 212
func FindWordII(board [][]byte, words []string)[]string{
	t:=TriesTree{}
	r:=[]string{}
	for i:=range words{
		t.InsertWord(words[i])
	}
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			dfs(i,j,board,&t,&r)
		}
	}
	return r
}
func dfs(i,j int ,board [][]byte,t *TriesTree,r *[]string){
	if i>len(board)||j>len(board[0])||i<0||j<0{
		return
	}
	if board[i][j]=='0'||t.next[board[i][j]-'a']==nil{
		return
	}
	directX:=[]int{0,-1,0,1}
	directY:=[]int{-1,0,1,0}
	tmp:=board[i][j]
	board[i][j]='0'
	t = t.next[tmp-'a']
	if t.isEnd{
		*r=append(*r,t.word)
		// one word use once
		t.word=""
		t.isEnd=false
	}
	for c:=0;c<4;c++{
		dfs(i+directX[c],j+directY[c],board,t,r)
	}
	board[i][j]=tmp
}
