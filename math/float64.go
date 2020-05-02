package math

import "math"

// Equals todo
func Equals(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

// Clamp todo
func Clamp(v, lo, hi float64) float64 {
	return math.Min(hi, math.Max(lo, v))
}

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float64) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	return f != f
}

// IsFinite reports whether f is neither NaN nor an infinity.
func IsFinite(f float64) bool {
	return !IsNaN(f - f)
}

// IsInf reports whether f is an infinity.
func IsInf(f float64) bool {
	return !IsNaN(f) && !IsFinite(f)
}
