package AbstractFactoryPattern

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

type Color interface {
	Fill() string
}

type Red struct {
}

func (Red) Fill() string {
	return "Red Fill"
}

type Green struct {
}

func (Green) Fill() string {
	return "Green Fill"
}

type Blue struct {
}

func (Blue) Fill() string {
	return "Blue Fill"
}

type AbstractFactory interface {
	GetColor(colorType string) Color
	GetShape(shapeType string) Shape
}

type ShapeFactory struct {
}

func (ShapeFactory) GetColor(colorType string) Color {
	return nil
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

type ColorFactory struct {
}

func (ColorFactory) GetShape(shapeType string) Shape {
	return nil
}
func (ColorFactory) GetColor(colorType string) Color {
	if colorType == "Red" {
		return new(Red)
	} else if colorType == "Green" {
		return new(Green)
	} else if colorType == "Blue" {
		return new(Blue)
	}
	return nil
}

type FactoryProducer struct {
}

func (FactoryProducer) GetFactory(choice string) AbstractFactory {
	if choice == "Shape" {
		return new(ShapeFactory)
	} else if choice == "Color" {
		return new(ColorFactory)
	}
	return nil
}
