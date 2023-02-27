package pizza_factory

import (
	"fmt"
)

// Pizzaer this interface defines the methods of Pizza
type Pizzaer interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

// Pizza pizza in reality
type Pizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (p *Pizza) Prepare() {
	fmt.Println("prepare", p.Name)
}

func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")

}

func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")

}

func (p *Pizza) Box() {
	fmt.Println("Place in official PizzaStore box")

}

type CheesePizza struct {
	Pizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{Pizza: Pizza{Name: "cheese pizza"}}
}

type PepperoniPizza struct {
	Pizza
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{Pizza{Name: "Pepperoni Pizza"}}
}

type ClamPizza struct {
	Pizza
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{
		Pizza: Pizza{Name: "Clam Pizza"},
	}
}

type VeggiePizza struct {
	Pizza
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{Pizza{Name: "Veggie Pizza"}}
}

type NYStyleCheesePizza struct {
	Pizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{Pizza: Pizza{
		Name:     "Chicago Style Deep Dish Cheese Pizza",
		Dough:    "Thin Crust Dough",
		Sauce:    "Marinara Sauce",
		Toppings: []string{"Grated Reggiano Cheese"},
	},
	}
}

type NYStyleClamPizza struct {
	Pizza
}

func NewNYStyleClamPizza() *NYStyleClamPizza {
	return &NYStyleClamPizza{Pizza{Name: "New York Style Clam Pizza"}}
}

type NYStyleVeggiePizza struct {
	Pizza
}

func NewNYStyleVeggiePizza() *NYStyleVeggiePizza {
	return &NYStyleVeggiePizza{Pizza{Name: "New York Style Veggie Pizza"}}
}

type NYStylePepperoniPizza struct {
	Pizza
}

func NewNYStylePepperoniPizza() *NYStylePepperoniPizza {
	return &NYStylePepperoniPizza{Pizza{Name: "New York Style Pepperoni Pizza"}}
}

type ChicagoStyleCheesePizza struct {
	Pizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{Pizza: Pizza{
		Name:     "Chicago Style Deep Dish Cheese Pizza",
		Dough:    "Extra Thick Crust Dough",
		Sauce:    "Plum Tomato Sauce",
		Toppings: []string{"Shredded Mozzarella Cheese"},
	},
	}
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
