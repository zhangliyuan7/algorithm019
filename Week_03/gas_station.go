package w3

//134
//good
func CanCompleteCircuit(gas []int, cost []int) int {
	left := 0   // 余量
	start := 0  // 初始选索引0作为起点
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		left += gas[i] - cost[i] // 累加每次的剩余量
		if left < 0 {    // 去不了下一站，0到i都不是起点
			start = i + 1 // 把i+1作为起点
			left = 0      // 余量归零
		}
	}
	if totalGas < totalCost { // 总油量不够，问题无解
		return -1
	}
	return start // 总加油>=总耗油，必然有解。当遍历结束时，最新的start指向成功的起点
}

