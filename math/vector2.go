package math

import "math"

// Vector2 todo
type Vector2 struct {
	X, Y float64
}

// IsZero todo
func (v Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// IsFinite todo
func (v Vector2) IsFinite() bool {
	return IsFinite(v.X) && IsFinite((v.Y))
}

// IsNormalized todo
func (v Vector2) IsNormalized() bool {
	return v.IsFinite() && math.Abs(v.Magnitude()-1) < 1e-4
}

// MagnitudeSquared todo
func (v Vector2) MagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Magnitude todo
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.MagnitudeSquared())
}

// Neg todo
func (v Vector2) Neg() Vector2 {
	return Vector2{-v.X, -v.Y}
}

// Add todo
func (v Vector2) Add(ov Vector2) Vector2 {
	return Vector2{v.X + ov.X, v.Y + ov.Y}
}

// Sub todo
func (v Vector2) Sub(ov Vector2) Vector2 {
	return Vector2{v.X - ov.X, v.Y - ov.Y}
}

// Mul todo
func (v Vector2) Mul(m float64) Vector2 {
	return Vector2{v.X * m, v.Y * m}
}

// Div todo
func (v Vector2) Div(m float64) Vector2 {
	return Vector2{v.X / m, v.Y / m}
}

// Dot todo
func (v Vector2) Dot(ov Vector2) float64 {
	return v.X*ov.X + v.Y*ov.Y
}

// Normalized todo
func (v Vector2) Normalized() Vector2 {
	m := v.MagnitudeSquared()
	if m > 0.0 {
		return v.Mul(m)
	}
	return Vector2{}
}

// Normalize todo
func (v *Vector2) Normalize() float64 {
	m := v.Magnitude()
	if m > 0.0 {
		n := v.Div(m)
		v.X = n.X
		v.Y = n.Y
	}
	return m
}

// Multiply todo
func (v Vector2) Multiply(ov Vector2) Vector2 {
	return Vector2{v.X * ov.X, v.Y * ov.Y}
}

// Minimum todo
func (v Vector2) Minimum(ov Vector2) Vector2 {
	return Vector2{math.Min(v.X, ov.X), math.Min(v.Y, ov.Y)}
}

// MinElement todo
func (v Vector2) MinElement() float64 {
	return math.Min(v.X, v.Y)
}

// Maximum todo
func (v Vector2) Maximum(ov Vector2) Vector2 {
	return Vector2{math.Max(v.X, ov.X), math.Max(v.Y, ov.Y)}
}

// MaxElement todo
func (v Vector2) MaxElement() float64 {
	return math.Max(v.X, v.Y)
}
