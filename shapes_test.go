package main

import (
	"math"
	"testing"
)


type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	Radius float64
}
type Shape interface{
	Area() float64
}

func Perimeter(width float64, height float64) float64{
	return 2*(width + height)
}

func TestPerimeter(t * testing.T){
	rectangle := Rectangle{10.0, 10.0} 
	got:= Perimeter(rectangle.Width, rectangle.Height)
	want:= 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T){

	checkArea := func (t * testing.T, shape Shape, want float64)  {
		t.Helper()
		got := shape.Area()
		if got != want{
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("retangle", func(t *testing.T) {
		rectangle := Rectangle{12.0,6.0} 
		// got:=rectangle.Area()
		// want:= 72.0
		checkArea(t,rectangle,72.0)
	})
	t.Run("cycle", func(t *testing.T) {
		circle := Circle{10} 
		// got:=circle.Area()
		// want:= 314.1592653589793
		checkArea(t,circle,314.1592653589793)
	
	})

	
}

func (r Rectangle) Area() float64{
	return r.Height *r.Width
}

func (c Circle) Area() float64{
	return c.Radius*c.Radius*math.Pi
}

// func Area(width float64, height float64) float64{
// 	return width * height
// }
// func Area(radius float64) float64{
// 	return radius * radius
// }