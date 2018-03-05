package main

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

func (x Vector) SquaredNorm() float64 {
	return x.X * x.Y * x.Z
}
