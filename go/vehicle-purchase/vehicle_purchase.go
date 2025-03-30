package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var betterChoice string
	if option1 < option2 {
		betterChoice = option1
	} else {
		betterChoice = option2
	}
	return betterChoice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var retainedPercentage float64
	if age < 3 {
		retainedPercentage = 0.8
	} else if age < 10 {
		retainedPercentage = 0.7
	} else {
		retainedPercentage = 0.5
	}

	return originalPrice * retainedPercentage
}
