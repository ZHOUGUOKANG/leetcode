package FactoryPattern

//interface
type Shape interface {
	Draw() string
}

type Rectangle struct {
}

func (Rectangle) Draw() string {
	return "Rectangle Draw"
}

type Square struct {
}

func (Square) Draw() string {
	return "Square Draw"
}

type Circle struct {
}

func (Circle) Draw() string {
	return "Circle Draw"
}

//factory
type ShapeFactory struct {
}

func (ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "Circle" {
		return new(Circle)
	} else if shapeType == "Rectangle" {
		return new(Rectangle)
	} else if shapeType == "Square" {
		return new(Square)
	}
	return nil
}
