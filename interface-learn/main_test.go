package interfacelearn

import (
	"testing"
)

func TestDetailShape(t *testing.T) {
	r := Rectangle{Length: 2.5}
	c := Circle{Radius: 2.5}
	t.Run("Test Circle Area", func(t *testing.T) {
		result := c.Area()
		if result != 19.625 {
			t.Errorf("Expected 19.625, got %f", result)
		}
	})
	t.Run("Test Rectangle Area", func(t *testing.T) {
		result := r.Area()
		if result != 6.25 {
			t.Errorf("Expected 6.25, got %f", result)
		}
	})
	t.Run("Test Rectangle Perimeter", func(t *testing.T) {
		result := r.Perimeter()
		if result != 10 {
			t.Errorf("Expected 10, got %f", result)
		}
	})
}

func TestShapeInfo(t *testing.T){
	r := Rectangle{Length: 2.5}
	
	got := getShapeInfo(r)
	expected := ShapeInfo{
		area: 6.25,
		perimeter: 10,
	}
	t.Logf("Expected %+v, result %+v", expected, got)
	result := got
	if result != expected {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}