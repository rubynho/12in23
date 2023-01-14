// Package weather is used to predict the weather.
package weather

// CurrentCondition variable stores the current weather condition.
var CurrentCondition string
// CurrentLocation variable stores the location where the weather forecast is done.
var CurrentLocation string

/*
Forecast predicts the weather using two variables:

city(string): the city where the weather is being predicted;
condition(string): the weathers condition for that city.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
