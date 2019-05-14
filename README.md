# tempconv [![Build Status](https://travis-ci.org/syscll/tempconv.svg?branch=master)](https://travis-ci.org/syscll/tempconv)

A simple package that performs the conversion of different temperature scales: `Celsius`, `Fahrenheit` and `Kelvin`.

## Usage

Simply call a specific conversion function and pass the correct scale:

```
CelsiusToFahrenheit(20)    // 68
CelsiusToKelvin(30)        // 303.2
FahrenheitToCelsius(55)    // 12.8
FahrenheitToKelvin(50)     // 283.2
KelvinToCelsius(300)       // 26.9
KelvinToFahrenheit(300)    // 80.3
```

Each type also implements the `String()` method which returns a formatted temperature:

```
fmt.Sprint(Celsius(12))    // 12°C
fmt.Sprint(Fahrenheit(55)) // 55°F
fmt.Sprint(Kelvin(300))    // 300K
```
