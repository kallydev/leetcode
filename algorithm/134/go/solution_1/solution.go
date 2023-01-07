package solution

func canCompleteCircuit(gas []int, cost []int) int {
	var totalTank, currentTank, startingStation int

	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]

		if currentTank < 0 {
			startingStation = i + 1
			currentTank = 0
		}
	}

	if totalTank >= 0 {
		return startingStation
	} else {
		return -1
	}
}
