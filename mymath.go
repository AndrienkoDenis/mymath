package mymath

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(1, x)
}

// Acos возвращает арккосинус x (в радианах).
func Acos(x float64) float64 {
    return math.Acos(x)
}

// Acosh возвращает гиперболический арккосинус x.
func Acosh(x float64) float64 {
    return math.Acosh(x)
}

// Asin возвращает арксинус x (в радианах).
func Asin(x float64) float64 {
    return math.Asin(x)
}

// Asinh возвращает гиперболический арксинус x.
func Asinh(x float64) float64 {
    return math.Asinh(x)
}

// Atan возвращает арктангенс x (в радианах).
func Atan(x float64) float64 {
    return math.Atan(x)
}

// Atanh возвращает гиперболический арктангенс x.
func Atanh(x float64) float64 {
    return math.Atanh(x)
}
