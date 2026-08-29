//Package weather provides tools to display current weather at a particular location.
package weather

var (
    //CurrentCondition represents current weather at a location.
	CurrentCondition string
    //CurrentLocation represents the location whose weather is displayed.
	CurrentLocation  string
)
//Forecast returns a string value that returns the weather condition of a location by accepting arguments.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}