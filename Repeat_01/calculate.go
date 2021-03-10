package Repeat_01

import "strconv"

//224

func calculate(s string) (ans int) {
	ops := []int{1}
	//只有加减法，所以直接括号符号顺应括号前符号，维护符号即可
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				// 多位数
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}

func calculateBAD(s string) int {
	sp:=[]byte{}
	for i,c:=range s{
		if c==' '{
			continue
		}else if c==')'{
			tmp:=[]byte{}
			for sp[len(sp)-1]!='('{
				tmp=append(tmp,sp[len(sp)-1])
				sp=sp[:len(sp)-1]
			}
			sp=sp[:len(sp)-1]
			sp=append(sp,byte(calF(tmp,-1)))
		}else{
			sp=append(sp,s[i])
		}
	}
	return calF(sp,1)
}

func calF(sp []byte,sign int)(r int){
	ln:=len(sp)
	if ln==0{
		return
	}
	if sign==1{
		r,_:=strconv.Atoi(string(sp[0]))
		for i:=1;i<ln-1;i++{
			if sp[i]=='+'{
				x,_:=strconv.Atoi(string(sp[i+1]))
				r+=x
			}else if sp[i]=='-'{
				x,_:=strconv.Atoi(string(sp[i+1]))
				r-=x
			}
		}
	}else{
		r,_:=strconv.Atoi(string(sp[ln-1]))
		for i:=ln-2;i>0;i--{
			if sp[i]=='+'{
				x,_:=strconv.Atoi(string(sp[i-1]))
				r+=x
			}else if sp[i]=='-'{
				x,_:=strconv.Atoi(string(sp[i-1]))
				r-=x
			}
		}
	}
	return
}