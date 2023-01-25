// Package weather provides tools to forecast the weather.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the current location to forecasting weather.
var CurrentLocation string

// Forecast returns a string with the forecast weather for the location,
// Forecast receives the strings city defining the location and
// condition defining the weather in the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
