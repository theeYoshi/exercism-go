package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var totalcost int
	var count int
	var remainder int

	remainder = carsCount

	for remainder >= 10 {
		remainder -= 10
		count++
	}

	totalcost = (count * 95000) + (remainder * 10000)
	fmt.Println(totalcost)
	return uint(totalcost)
}
