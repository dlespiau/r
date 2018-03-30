package main

// Point is a point on a plane.
type Point struct {
	X, Y float64
}

// point is the internal representation of a point.
type point struct {
	X, Y int32
}
