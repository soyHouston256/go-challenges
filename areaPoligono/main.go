package main

import "github.com/soyHouston256/go-challenges/areaPoligono/algorithm"

func main() {
	algorithm.Area(algorithm.Triangle{Base: 10, Height: 10})
	algorithm.Area(algorithm.Square{Base: 10})
	algorithm.Area(algorithm.Rectangle{Base: 10, Height: 5})
}
