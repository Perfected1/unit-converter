package converter

// -----------------------------
// Weight Conversions
// -----------------------------

// KgToLbs converts kilograms to pounds
func KgToLbs(kg float64) float64 {
	return kg * 2.20462
}

// LbsToKg converts pounds to kilograms
func LbsToKg(lbs float64) float64 {
	return lbs / 2.20462
}

// -----------------------------
// Distance Conversions
// -----------------------------

// KmToMiles converts kilometers to miles
func KmToMiles(km float64) float64 {
	return km * 0.621371
}

// MilesToKm converts miles to kilometers
func MilesToKm(miles float64) float64 {
	return miles / 0.621371
}

// -----------------------------
// Volume Conversions
// -----------------------------

// LitersToGallons converts liters to gallons
func LitersToGallons(liters float64) float64 {
	return liters * 0.264172
}

// GallonsToLiters converts gallons to liters
func GallonsToLiters(gallons float64) float64 {
	return gallons / 0.264172
}