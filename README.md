# Temperature [![Build Status](https://travis-ci.org/danbondd/temperature.svg?branch=master)](https://travis-ci.org/danbondd/temperature)

A simple package that allows the conversion from Celsius to Fahrenheit and vice-versa.

## Usage

The package currently supports two different temperature scales: `Celsius` and `Fahrenheit`

Simply call a specific conversion function and pass the correct scale:

```
fahrenheit := converter.CelsiusToFahrenheit(20) // 68
celsius := converter.FahrenheitToCelsius(55) // 12.8
```

Additionally, you can call the `String()` function to return a formatted version:

```
celsius := converter.Celsius(12).String() // 12°C
fahrenheit := converter.Fahrenheit(55).String() // 55°F
```

## To-Do

- Add Kelvin conversion
