package Repeat_01

//132
func minCut(s string) int {
	if len(s)<=1{
		return 0
	}
	ln:=len(s)
	g:=make([][]bool,len(s))
	for i:=range g{
		g[i]=make([]bool,ln)
		for j:=range g[i]{
			g[i][j]=true
		}
	}
	// 回文dp初始化
	for i:=ln-1;i>=0;i--{
		for j:=i+1;j<ln;j++{
			// 首等于尾，且包含的子串是回文
			// 当前即为回文
			// 这里蛮吊的
			g[i][j]=s[i]==s[j]&&g[i+1][j-1]
		}
	}
	// 结果集
	f:=make([]int,ln)
	// 循环从j到i  ，f[i]= min(f[j]+1,f[i])[0<=j<i] => g[j][i]==true
	for i:=range f{
		// 0-i 是回文的 直接跳过 不用切割
		if g[0][i]{
			f[i]=0
			continue
		}
		// f[i]最多也就是len s[i] 的数量
		f[i]=i
		// 从0-j的部分 递推 到 i ，因为要使得 f[i]=f[j]+1
		// 从大到小的话 f[j]还没计算出来 所以为0，结果就会不对
		for j:=0;j<i;j++{
			// j+1 到i是个回文，那么f[i]=f[j]+1
			// 如果g[j+1(0<=j<i)][i]一直不是回文，即最差的情况：一直没有回文出现
			// 直到i-1=j 剩余的单个字母s[i]一定是回文，所以f[i]=min(f[i-1]+1,f[i])
			// 故至少会有一次执行进入第二级别判断，而不会一次都不执行
			if g[j+1][i]{
				// 找到当前最小的，最差 = i-1+1
				if f[j]+1<f[i]{
					f[i]=f[j]+1
				}
			}
		}
	}
	return f[ln-1]
}

func minCutII(s string) int {
	if len(s)<1{
		return 0
	}
	ln:=len(s)
	g:=make([][]bool,ln)
	for i:=range g{
		g[i]=make([]bool,ln)
	}
	//1 all prom
	for i:=ln-1;i>=0;i--{
		for j:=i+1;j<ln;j++{
			g[i][j]=s[i]==s[j]&&g[i+1][j-1]
		}
	}
	f:=make([]int,ln)
	for i:=range f{
		if g[0][i]{
			f[i]=0
			continue
		}
		// init f[i]
		f[i]=i
		// for i for j-i dp
		for j:=0;j<i;j++{
			// j+1 - i
			if g[j+1][i]{
				f[i]=min(f[i],f[j]+1)
			}
		}
	}
	return f[ln-1]
}

func minCutIII(s string) int {
	ln:=len(s)
	if ln<=1{
		return 0
	}
	//g[i][j]
	g:=make([][]bool,ln)
	for i:=range g{
		g[i]=make([]bool,ln)
		for j:=range g[i]{
			g[i][j]=true
		}
	}
	for i:=ln-1;i>=0;i--{
		for j:=i+1;j<ln;j++{
			g[i][j]=s[i]==s[j]&&g[i+1][j-1]
		}
	}
	f:=make([]int,ln)
	for i:=0;i<ln;i++{
		if g[0][i]{
			f[i]=0
			continue
		}
		f[i]=i
		for j:=0;j<i;j++{
			if g[j+1][i]{
				if f[j]+1<f[i]{
					f[i]=f[j]+1
				}
			}
		}
	}
	return f[ln-1]
}







