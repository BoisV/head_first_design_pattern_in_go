package caffeine_beverage

import "fmt"

type ICaffeineBeverage interface {
	Brew()
	AddCondiments()
}

type CaffeineBeverage struct {
	ICaffeineBeverage
}

func NewCaffeineBeverage(ICaffeineBeverage ICaffeineBeverage) *CaffeineBeverage {
	return &CaffeineBeverage{ICaffeineBeverage: ICaffeineBeverage}
}

func (b *CaffeineBeverage) PrepareRecipe() {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}

func (b *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}
