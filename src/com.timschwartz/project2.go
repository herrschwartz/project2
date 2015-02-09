package main

import "fmt"
import "math"

type Bundle struct {
	Root1 float64
	Root2 float64
}

func main() {
	var x, y, z float64 = 3, 30, 3
	var pos, _ = quadratic(x, y, z)

	var bundle Bundle
	bundle.Root1 = pos
	bundle.Root2 = 0

	fmt.Println(pos)

}

func quadratic(a, b, c float64) (float64, float64) {
	var temp float64 = (b * b) - 4*a*c
	var x1 float64 = (-b + math.Sqrt(temp)) / (2 * a)
	var x2 float64 = (-b - math.Sqrt(temp)) / (2 * a)
	return x1, x2
}
