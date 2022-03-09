package main

import (
	"math"
)

type NewSquare struct {
	a float64
}

type NewCircle struct {
	r float64
}

func (square *NewSquare) Area() float64 {
	squareArea := square.a * square.a
	return squareArea
}

func (circle *NewCircle) Area() float64 {
	circleArea := math.Pi * math.Pow(circle.r, 2)
	return circleArea
}

type Shape interface {
	Area()
}

type Shapes []Shape

func (Shapes) largestArea() float64 {
	var slice []Shape
	biggestFigure := slice[0]
	for _,i := range slice {
		if biggestFigure < i {
			biggestFigure = float64(i)
		}
	}
	return biggestFigure
}

func main(){
	square1 := NewSquare{5}
	square2 := NewSquare{4}
	circle := NewCircle{8}
	var slice Shapes
	slice = append(slice,)
}
