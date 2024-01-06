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

type Triangle struct{
		Base float64
		Height float64
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

// func TestArea(t *testing.T){

// 	checkArea := func (t * testing.T, shape Shape, want float64)  {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want{
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}
// 	t.Run("retangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0,6.0} 
// 		// got:=rectangle.Area()
// 		// want:= 72.0
// 		checkArea(t,rectangle,72.0)
// 	})
// 	t.Run("cycle", func(t *testing.T) {
// 		circle := Circle{10} 
// 		// got:=circle.Area()
// 		// want:= 314.1592653589793
// 		checkArea(t,circle,314.1592653589793)
	
// 	})

	
// }


func TestArea(t *testing.T){
	areaTests:= []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12,6}, hasArea:72},
		{name: "Circle", shape: Circle{10}, hasArea:314.1592653589793},
		{name: "Triangle",shape: Triangle{12,6}, hasArea:36.0},
	}

	for _, tt:= range areaTests{
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea{
				t.Errorf("%#v got %.2f want %.2f",tt.shape, got, tt.hasArea)
			}
		})
		
	}
}



func (r Rectangle) Area() float64{
	return r.Height *r.Width
}

func (c Circle) Area() float64{
	return c.Radius*c.Radius*math.Pi
}

func (c Triangle) Area() float64{
	return (c.Base * c.Height) * 0.5
}

// func Area(width float64, height float64) float64{
// 	return width * height
// }
// func Area(radius float64) float64{
// 	return radius * radius
// }