package main

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (x Vector) Add(y Vector) Vector {
	return Vector{
		X: x.X + y.X,
		Y: x.Y + y.Y,
		Z: x.Z + y.Z,
	}
}

func (x Vector) Sub(y Vector) Vector {
	return Vector{
		X: x.X - y.X,
		Y: x.Y - y.Y,
		Z: x.Z - y.Z,
	}
}

func (x Vector) ScalarMul(y float64) Vector {
	return Vector{
		X: x.X * y,
		Y: x.Y * y,
		Z: x.Z * y,
	}
}

func (x Vector) SquaredNorm() float64 {
	return x.X * x.Y * x.Z
}

func (x Vector) Norm() float64 {
	return math.Sqrt(x.SquaredNorm())
}
