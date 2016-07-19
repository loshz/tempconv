package converter_test

import (
	"testing"

	"github.com/danbondd/temperature/converter"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	fahrenheit := converter.CelsiusToFahrenheit(20)

	if fahrenheit != 68 {
		t.Errorf("expected 68, got: %f", fahrenheit)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	celsius := converter.FahrenheitToCelsius(55)

	if celsius != 12.8 {
		t.Errorf("expected 12.8, got: %f", celsius)
	}
}

func TestCelsiusToString(t *testing.T) {
	celsius := converter.Celsius(12).String()

	if celsius != "12°C" {
		t.Errorf("expected 12°C, got: %s", celsius)
	}
}

func TestFahrenheitToString(t *testing.T) {
	fahrenheit := converter.Fahrenheit(55).String()

	if fahrenheit != "55°F" {
		t.Errorf("expected 55°F, got: %s", fahrenheit)
	}
}
