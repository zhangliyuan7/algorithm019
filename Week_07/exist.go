package Week_07

// 79
// 单词搜索
func Exist(board [][]byte, word string) bool {
	if len(board)==0||len(board[0])==0{
		return false
	}
	for i:=0;i<len(board);i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] &&
				dfs(board, i, j, word){
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte,i,j int ,word string)bool{
	if i>=len(board)||j>=len(board[0])||i<0||j<0{
		return false
	}
	if word[0]!=board[i][j]{
		return false
	}
	if len(word)==1&&word==string(board[i][j]){
		return true
	}
	// east south west north
	tmp:=board[i][j]
	board[i][j]='0'
	var dx = []int{-1,0,1,0}
	var dy = []int{0,-1,0,1}
	for x:=0;x<4;x++{
		if dfs(board,i+dx[x],j+dy[x],word[1:]){
			return true
		}
	}
	board[i][j]=tmp
	return false
}