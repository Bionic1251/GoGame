package server

import (
	"math"
)

type Coordinate struct {
	X float64
	Y float64
}

type Location struct {
	Coordinate
	Radius float64
}

func (l Location) Contains(c Coordinate) bool {
	if l.DistanceFrom(c) <= l.Radius {
		return true
	}
	return false
}

func (c Coordinate) DistanceFrom(o Coordinate) float64 {
	dx := math.Pow(o.X-c.X, 2)
	dy := math.Pow(o.Y-c.Y, 2)
	d := math.Sqrt(dx + dy)

	return d
}

type Area []Container

type Container interface {
	Contains(Coordinate) bool
}

type Cafeteria struct {
	Location
}
