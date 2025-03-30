// Package weather does weather things.
package weather

// CurrentCondition holds the description of the current weather condition for a specified location.
var CurrentCondition string

// CurrentLocation holds the name of the location for which the current weather condition is being specified.
var CurrentLocation string

// Forecast updates and returns the current weather condition for a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
