package greetings

import "math"

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsNormal(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	// https://go.dev/tour/methods/4
	// https://go.dev/tour/methods/8
	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
	v.X *= f
	v.Y *= f
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}
