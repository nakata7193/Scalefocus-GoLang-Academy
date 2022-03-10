package main

import (
	"fmt"
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
	Area() float64
}

type Shapes []Shape

func (shapes Shapes) largestArea() float64 {
	biggestFigure := shapes[0].Area()
	for _,i := range shapes {
		if biggestFigure < float64(i.Area()) {
			biggestFigure = float64(i.Area())
		}
	}
	return biggestFigure
}

func main() {
	square1 := NewSquare{5}
	square2 := NewSquare{4}
	circle := NewCircle{8}
	var slice Shapes
	slice = append(slice,&square1)
	slice = append(slice,&square2)
	slice = append(slice,&circle)
	fmt.Print(slice.largestArea())
}
