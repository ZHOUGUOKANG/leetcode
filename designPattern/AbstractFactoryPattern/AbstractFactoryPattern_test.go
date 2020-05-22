package AbstractFactoryPattern

import (
	"testing"
)

func Test_AbstractFactoryPattern_Rectangle_Draw(t *testing.T) {
	result := FactoryProducer{}.GetFactory("Shape").GetShape("Rectangle").Draw()
	if result != "Rectangle Draw" {
		t.Error()
	}
}
func Test_AbstractFactoryPattern_Circle_Draw(t *testing.T) {
	result := FactoryProducer{}.GetFactory("Shape").GetShape("Circle").Draw()
	if result != "Circle Draw" {
		t.Error()
	}
}

func Test_AbstractFactoryPattern_Blue_Fill(t *testing.T) {
	result := FactoryProducer{}.GetFactory("Color").GetColor("Blue").Fill()
	if result != "Blue Fill" {
		t.Error()
	}
}

func Test_AbstractFactoryPattern_Red_Fill(t *testing.T) {
	result := FactoryProducer{}.GetFactory("Color").GetColor("Red").Fill()
	if result != "Red Fill" {
		t.Error()
	}
}
