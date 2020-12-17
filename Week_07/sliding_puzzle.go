package Week_07

import "fmt"

// 773 滑动谜题
// bfs
// 0 1 2
// 3 4 5
func SlidingPuzzle(board [][]int) int {
	moves:=make(map[int][]int)
	moves[0]=[]int{1,3}
	moves[1]=[]int{0,2,4}
	moves[2]=[]int{1,5}
	moves[3]=[]int{0,4}
	moves[4]=[]int{1,3,5}
	moves[5]=[]int{2,4}
	s:=""
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[i]);j++{
			s+=fmt.Sprintf("%d",board[i][j])
		}
	}

	endState:="123450"
	if s==endState{
		return 0
	}
	startQ:=make(map[string]bool)
	endQ:=make(map[string]bool)
	endQ[endState]=true
	startQ[s]=true
	count:=1
	visited:=make(map[string]bool)
	for len(startQ)!=0{
		tmpQ:=make(map[string]bool)
		for sortStr:=range startQ{
			zid:=zeroIndex(sortStr)
			tmpMove:=moves[zid]
			for zin:=0;zin<len(tmpMove);zin++{
				new_string:=[]byte(sortStr)
				new_string[zid],new_string[tmpMove[zin]]=new_string[tmpMove[zin]],new_string[zid]
				if _,ok:=endQ[string(new_string)];ok{
					return count
				}
				if _,ok:=visited[string(new_string)];!ok{
					tmpQ[string(new_string)]=true
					visited[string(new_string)]=true
				}
			}
		}
		startQ=tmpQ
		if len(startQ)>len(endQ){
			startQ,endQ=endQ,startQ
		}
		count++
	}
	return -1
}
func zeroIndex(s string)int{
	for i,c:=range s{
		if string(c)=="0"{
			return i
		}
	}
	return 0
}