package Week_06

//649
func PredictPartyVictory(senate string) string {
	//values are their indexes
	r:=[]int{}
	d:=[]int{}
	for i:=range senate{
		if string(senate[i])=="R"{
			r=append(r,i)
		}else{
			d=append(d,i)
		}
	}
	for len(r)!=0&&len(d)!=0{
		fir_r:=r[0]
		fir_d:=d[0]
		if fir_r<fir_d{
			r=append(r,fir_r+len(senate))
		}else{
			d=append(d,fir_d+len(senate))
		}
		r=r[1:]
		d=d[1:]
	}
	if len(r)>0{
		return "Radiant"
	}
	return "Dire"
}