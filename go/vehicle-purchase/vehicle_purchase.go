package purchase

const (
	SmallDepreciation  = 0.8
	MediumDepreciation = 0.7
	BigDepreciation    = 0.5
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var rightChoice string

	if option1 < option2 {
		rightChoice = option1
	} else {
		rightChoice = option2
	}

	return rightChoice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * SmallDepreciation
	} else if age > 3 && age < 10 {
		return originalPrice * MediumDepreciation
	} else {
		return originalPrice * BigDepreciation
	}
}
