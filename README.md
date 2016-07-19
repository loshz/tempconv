# Temperature [![Build Status](https://travis-ci.org/danbondd/temperature.svg?branch=master)](https://travis-ci.org/danbondd/temperature)

A simple package that allows the conversion from Celsius to Fahrenheit and vice-versa.

## Usage

The package currently supports three different temperature scales: `Celsius`, `Fahrenheit` and `Kelvin`

Simply call a specific conversion function and pass the correct scale:

```
converter.CelsiusToFahrenheit(20) // 68
converter.CelsiusToKelvin(30) // 303.2
converter.FahrenheitToCelsius(55) // 12.8
converter.FahrenheitToKelvin(50) // 283.2
converter.KelvinToCelsius(300) // 26.9
converter.KelvinToFahrenheit(300) // 80.3
```

Each type also implements the `String()` method which returns a formatted temperature:

```
fmt.Sprintf("%s", converter.Celsius(12)) // 12°C
fmt.Sprintf("%s", converter.Fahrenheit(55)) // 55°F
fmt.Sprintf("%s", converter.Kelvin(300)) // 300K
```
