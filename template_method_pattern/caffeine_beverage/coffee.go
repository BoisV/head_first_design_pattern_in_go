package caffeine_beverage

import "fmt"

type Coffee struct {
	CaffeineBeverage
}

func NewCoffee() *Coffee {
	return &Coffee{}
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
