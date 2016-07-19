package converter

import (
	"fmt"
	"math"
)

// Celsius is a scale and unit of measurement for temperature.
type Celsius float64

// Fahrenheit is a scale and unit of measurement for temperature.
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// CelsiusToFahrenheit converts celsius to fahrenheit.
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	value := round(float64((c * 9 / 5) + 32))
	return Fahrenheit(value)
}

// FahrenheitToCelsius converts fahrenheit to celsius.
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	value := round(float64((f - 32) * 5 / 9))
	return Celsius(value)
}

func round(val float64) float64 {
	var round float64
	digit := 10 * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / 10
}
