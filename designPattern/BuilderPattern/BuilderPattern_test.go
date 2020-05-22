package BuilderPattern

import (
	"fmt"
	"testing"
)

func Test_BuilderPattern(t *testing.T) {
	mealBuilder := new(MealBuilder)
	fmt.Println(mealBuilder.PrepareVegMeal().GetCost())
	fmt.Println(mealBuilder.PrepareVegMeal().ShowItems())
	fmt.Println(mealBuilder.PrepareNonVegMeal().GetCost())
	fmt.Println(mealBuilder.PrepareNonVegMeal().ShowItems())
}
