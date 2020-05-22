package BuilderPattern

type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}
type Packing interface {
	Pack() string
}

type Wrapper struct {
}

func (Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct {
}

func (Bottle) Pack() string {
	return "Bottle"
}

type Burger interface {
	Name() string
	Packing() Packing
	Price() float64
}

type ColdDrink interface {
	Name() string
	Packing() Packing
	Price() float64
}

type VegBurger struct {
}

func (*VegBurger) Name() string {
	return "VegBurger"
}
func (*VegBurger) Price() float64 {
	return 25.0
}
func (*VegBurger) Packing() Packing {
	return new(Wrapper)
}

type ChickenBurger struct {
}

func (*ChickenBurger) Name() string {
	return "ChickenBurger"
}
func (*ChickenBurger) Price() float64 {
	return 50.5
}
func (*ChickenBurger) Packing() Packing {
	return new(Wrapper)
}

type Coke struct {
	ColdDrink
}

func (*Coke) Name() string {
	return "Coke"
}
func (*Coke) Price() float64 {
	return 30.0
}
func (*Coke) Packing() Packing {
	return new(Bottle)
}

type Pepsi struct {
	ColdDrink
}

func (*Pepsi) Name() string {
	return "Pepsi"
}
func (*Pepsi) Price() float64 {
	return 35.0
}
func (*Pepsi) Packing() Packing {
	return new(Bottle)
}

type Meal struct {
	items []Item
}

func (meal *Meal) AddItem(item Item) {
	meal.items = append(meal.items, item)
}

func (meal *Meal) GetCost() float64 {
	cost := 0.0
	for _, item := range meal.items {
		cost += item.Price()
	}
	return cost
}

func (meal *Meal) ShowItems() string {
	names := ""
	for _, item := range meal.items {
		names += "item:" + item.Name() + "Packing:" + item.Packing().Pack() + "\n"
	}
	return names
}

type MealBuilder struct {
}

func (MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(VegBurger))
	meal.AddItem(new(Coke))
	return meal
}

func (MealBuilder) PrepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(Pepsi))
	return meal
}
