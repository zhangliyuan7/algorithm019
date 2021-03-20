package Repeat_01

//1603
type ParkingSystem struct {
	space [4]int
}

func ConstructorParkingSystem(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[4]int{0,big,medium,small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.space[carType]>0{
		this.space[carType]-=1
		return true
	}
	return false
}

//位运算
type ParkingSystemI struct {
	count int
}

func ConstructorParkingSystemI(big int, medium int, small int) ParkingSystemI {
	return ParkingSystemI{big<<20|medium<<10|small}
}

func (this *ParkingSystemI) AddCar(carType int) bool {
	bit:=carType*(3-carType)*10
	if (this.count>>bit)&0b1111111111<=0{
		return false
	}else{
		this.count-=1<<bit
	}
	return true
}




/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */