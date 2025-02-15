// Package weather is just for learning purpose.
package weather

// CurrentCondition represents something :).
var CurrentCondition string

// CurrentLocation also represents something else ;).
var CurrentLocation string

// Forecast returns a string value as you can see and this comment have no meaning.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
