package cars

const AnHourInMinutes = 60
const ProductionCostPerCar = 10_000
const ProductionCostPerTenCars = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := int(CalculateWorkingCarsPerHour(productionRate, successRate))

	return workingCarsPerHour / AnHourInMinutes
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTenCars := carsCount / 10
	remainingCars := carsCount % 10

	return uint(groupOfTenCars * ProductionCostPerTenCars + remainingCars * ProductionCostPerCar)
}
