package starbuzz_coffee

type Beverageer interface {
	Cost() float64
	GetDescription() string
}

type Beverage struct {
	Description string
	Price       float64 // 价格
	Count       float64 // 折扣

}

func (b *Beverage) Cost() float64 {
	return b.Price * b.Count
}

func (b *Beverage) GetDescription() string {
	return b.Description
}

type DarkRoast struct {
	Beverage
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{Beverage{
		Description: "dark roast",
		Price:       0.99,
		Count:       1,
	}}
}

type HouseBlend struct {
	Beverage
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{Beverage{
		Description: "house blend",
		Price:       0.89,
		Count:       1,
	}}
}

type Espresso struct {
	Beverage
}

func NewEspresso() *Espresso {
	return &Espresso{Beverage{
		Description: "espresso",
		Price:       1.99,
		Count:       1,
	}}
}

type Decaf struct {
	Beverage
}

func NewDecaf() *Decaf {
	return &Decaf{Beverage{
		Description: "decaf",
		Price:       1.05,
		Count:       1,
	}}
}
