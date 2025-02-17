package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		option1 = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", option1)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	rate := 0.7
	switch {
	case age < 3:
		rate = 0.8
	case age >= 10:
		rate = 0.5
	}

	return originalPrice * rate
}
