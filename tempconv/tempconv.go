package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// AbsoluteZeroC representa o zero absoluto em Celsius
	AbsoluteZeroC Celsius = -273.15
	// AbsoluteZeroK representa o zero absoluto em Kelvin
	AbsoluteZeroK Kelvin = 0

	// FreezingC representa o ponto de congelamento da água em Celsius
	FreezingC Celsius = 0
	// BoilingC representa o ponto de ebulição da água em Celsius
	BoilingC Celsius = 100
)

// String imprime uma temperatura em Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String imprime uma temperatura em Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// String imprime uma temperatura em Kelvin
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

// CToF converte Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converte Celsius para Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converte Kelvin para Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
