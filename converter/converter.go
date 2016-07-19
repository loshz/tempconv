package converter

import (
	"fmt"
	"math"
)

// Celsius is a scale and unit of measurement for temperature.
type Celsius float64

// Fahrenheit is a scale and unit of measurement for temperature.
type Fahrenheit float64

// Kelvin is a scale and unit of measurement for temperature.
type Kelvin float64

const (
	absoluteZeroC = 273.15
	absoluteZeroF = 459.67
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

// CelsiusToFahrenheit converts celsius to fahrenheit.
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	value := round(float64((c * 9 / 5) + 32))
	return Fahrenheit(value)
}

// CelsiusToKelvin converts celsius to kelvin.
func CelsiusToKelvin(c Celsius) Kelvin {
	value := round(float64(c + absoluteZeroC))
	return Kelvin(value)
}

// FahrenheitToCelsius converts fahrenheit to celsius.
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	value := round(float64((f - 32) * 5 / 9))
	return Celsius(value)
}

// FahrenheitToKelvin converts fahrenheit to kelvin.
func FahrenheitToKelvin(f Fahrenheit) Kelvin {
	value := round(float64((f + absoluteZeroF) * 5 / 9))
	return Kelvin(value)
}

// KelvinToCelcius converts kelvin to celsius.
func KelvinToCelcius(k Kelvin) Celsius {
	value := round(float64(k - absoluteZeroC))
	return Celsius(value)
}

// KelvinToFahrenheit converts kelvin to fahrenheit.
func KelvinToFahrenheit(k Kelvin) Fahrenheit {
	value := round(float64((k * 9 / 5) - absoluteZeroF))
	return Fahrenheit(value)
}

// Round a value to 1 decimal place.
func round(val float64) float64 {
	var round float64
	digit := 10 * val
	if _, div := math.Modf(digit); div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / 10
}
