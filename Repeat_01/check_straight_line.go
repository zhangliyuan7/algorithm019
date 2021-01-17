package Repeat_01

//1232
//y=kx+b
//y2=kx2+b
//y2-y1=k(x2-x1)
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates)<2{
		return true
	}
	leftx,lefty:=coordinates[0][0],coordinates[0][1]
	//去掉直线方程中的b ，平移经过原点
	for i:=range coordinates{
		coordinates[i][0]-=leftx
		coordinates[i][1]-=lefty
	}
	secPx:=coordinates[1][0]
	secPy:=coordinates[1][1]
	//y1/x1=y2/x2
	//y1x2=y2x1
	//y1x2-y2x1=0

	//从第三个点向后计算每个点与第二个点所在的直线是否经过原点
	for i:=2;i<len(coordinates);i++{
		if coordinates[i][0]*secPy-coordinates[i][1]*secPx!=0{
			return false
		}
	}
	return true
}