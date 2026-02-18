// Package weather provides forecast of the current location and weather.
package weather


var (
    // CurrentCondition stores a state name of the weather at the moment.
	CurrentCondition string
    // CurrentLocation is a name of the city right now.
	CurrentLocation  string
)

// Forecast takes city and condition as  strings and return prediction of the weather, also in form of string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
