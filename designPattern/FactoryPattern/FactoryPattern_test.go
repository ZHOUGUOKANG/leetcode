package FactoryPattern

import (
	"testing"
)

func Test_FactoryPattern_Rectangle_Draw(t *testing.T) {
	result := ShapeFactory{}.GetShape("Rectangle").Draw()
	if result != "Rectangle Draw" {
		t.Error()
	}
}

func Test_FactoryPattern_Circle_Draw(t *testing.T) {
	result := ShapeFactory{}.GetShape("Circle").Draw()
	if result != "Circle Draw" {
		t.Error()
	}
}
