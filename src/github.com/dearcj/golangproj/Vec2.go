package main

import (
	"math"
)

type Vec2 [2]float32
type Vec2f64 [2]float64

type IntVec2 [2]int

func (a *Vec2) set(v *Vec2) {
	a[0] = v[0]
	a[1] = v[1]
}

func (a *Vec2f64) to32() *Vec2 {
	return &Vec2{float32(a[0]), float32(a[1])}
}

func (a *Vec2f64) len() float64 {
	return math.Sqrt(a[0]*a[0] + a[1]*a[1])
}

func (a *Vec2) to64() *Vec2f64 {
	return &Vec2f64{float64(a[0]), float64(a[1])}
}

func (a *Vec2) copy() *Vec2 {
	return &Vec2{a[0], a[1]}
}

func (a *Vec2) rotate(angle float32) Vec2 {
	var ca = float32(math.Cos(float64(angle)))
	var sa = float32(math.Sin(float64(angle)))
	return Vec2{ca*a[0] - sa*a[1], sa*a[0] + ca*a[1]}
}

func sqdist(a *Vec2, b *Vec2) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return float64(dx*dx + dy*dy)
}

func (a *Vec2) len() float32 {
	return float32(math.Sqrt(float64(a[0]*a[0] + a[1]*a[1])))
}

func dist(a *Vec2, b *Vec2) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func dists(a *Vec2, b *Vec2) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return float64(dx*dx + dy*dy)
}
