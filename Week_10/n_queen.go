package Week_10

import (
	"math/bits"
)

//51
var solution [][]string
func solveNQueens(n int) [][]string {
	solution =[][]string{}
	queens:=make([]int,n)
	for i:=range queens{
		queens[i]=-1
	}
	solve(queens,n,0,0,0,0)
	return solution
}

func solve(queens []int,n,row,columns,lefts,rights int){
	if n==row{
		board:=generateBoard(queens,n)
		solution=append(solution,board)
		return
	}
	available:=((1<<n)-1)&(^(columns | lefts | rights))
	//look at here ，it's "FOR" LOOP
	for available!=0{
		// 获得该数字 仅包含其二进制最后一个1的值
		//(12)&(-12) = 8
		// 1100&(-(1100)) == 1100 & 0100 => 0100 => 8
		position:=available&(-available)
		//fmt.Println(position)
		available=available&(available-1)
		// 返回(position-1) 的二进制包含1的个数
		// 作为索引，因为position是最后一个1的位置（仅包含一个1 ，比如0100 ），-1（0011 ） 计算1的个数，即为其索引位置
		column:=bits.OnesCount(uint(position-1))
		//fmt.Println(column)
		// 倒序索引位置
		queens[row]=column
		solve(queens,n,row+1,columns|position,(lefts|position)<<1,(rights|position)>>1 )
		queens[row]=-1
	}
}

func generateBoard(queens []int,n int )[]string{
	board:=[]string{}
	for i:=0;i<n;i++{
		row:=make([]byte,n)
		for j:=0;j<n;j++{
			row[j]='.'
		}
		row[queens[i]]='Q'
		board=append(board,string(row))
	}
	return board
}