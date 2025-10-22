// Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

var (
	// CurrentCondition holds the current weather condition for the location of interest.
	CurrentCondition string

	// CurrentLocation holds the location for which the current weather condition is about.
	CurrentLocation string
)

// Forecast takes a city and a condition as arguments and returns a weather forecast of the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
