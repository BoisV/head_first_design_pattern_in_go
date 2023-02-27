package ingredient_facttory

import (
	"fmt"
)

// Pizzaer this interface defines the methods of Pizza
type Pizzaer interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

// Pizza pizza in reality
type Pizza struct {
	Name      string
	Dough     Dougher
	Sauce     Saucer
	Cheese    Cheeser
	Pepperoni Pepperonier
	Clam      Clamer
	Veggies   []Veggieer
}

func (p *Pizza) Prepare() {
	//TODO implement me
	panic("implement me")
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

func (p *Pizza) GetName() string {
	return p.Name
}

// CheesePizza 起司披萨
type CheesePizza struct {
	Pizza
	PizzaIngredientFactory PizzaIngredientFactorier
}

func NewCheesePizza(factory PizzaIngredientFactorier) *CheesePizza {
	return &CheesePizza{Pizza: Pizza{Name: "cheese pizza"}, PizzaIngredientFactory: factory}
}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing", c.Name)
	c.Dough = c.PizzaIngredientFactory.CreateDough()
	c.Sauce = c.PizzaIngredientFactory.CreateSauce()
	c.Cheese = c.PizzaIngredientFactory.CreateCheese()
}

// PepperoniPizza 帕拉尼披萨
type PepperoniPizza struct {
	Pizza
	PizzaIngredientFactory PizzaIngredientFactorier
}

func NewPepperoniPizza(factorier PizzaIngredientFactorier) *PepperoniPizza {
	return &PepperoniPizza{Pizza{Name: "Pepperoni Pizza"}, factorier}
}

// ClamPizza 蛤蜊披萨
type ClamPizza struct {
	Pizza
	PizzaIngredientFactory PizzaIngredientFactorier
}

func NewClamPizza(factory PizzaIngredientFactorier) *ClamPizza {
	return &ClamPizza{Pizza: Pizza{Name: "Clam Pizza"}, PizzaIngredientFactory: factory}
}

func (c *ClamPizza) Prepare() {
	fmt.Println("Preparing", c.Name)
	c.Dough = c.PizzaIngredientFactory.CreateDough()
	c.Sauce = c.PizzaIngredientFactory.CreateSauce()
	c.Cheese = c.PizzaIngredientFactory.CreateCheese()
	c.Clam = c.PizzaIngredientFactory.CreateClam()
}

type VeggiePizza struct {
	Pizza
	PizzaIngredientFactory PizzaIngredientFactorier
}

func NewVeggiePizza(factorier PizzaIngredientFactorier) *VeggiePizza {
	return &VeggiePizza{Pizza{Name: "Veggie Pizza"}, factorier}
}

type NYStyleCheesePizza struct {
	Pizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{Pizza: Pizza{
		Name:      "Chicago Style Deep Dish Cheese Pizza",
		Dough:     &ThinCrustDough{},
		Sauce:     &MarinaraSauce{},
		Cheese:    &ReggianoCheese{},
		Pepperoni: &SlicePepperoni{},
		Clam:      &FreshClam{},
		Veggies:   []Veggieer{&Garlic{}, &Onion{}, &Mushroom{}, &RedPepper{}},
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
		Name:      "Chicago Style Deep Dish Cheese Pizza",
		Dough:     &ThinCrustDough{},
		Sauce:     &MarinaraSauce{},
		Cheese:    &ReggianoCheese{},
		Pepperoni: &SlicePepperoni{},
		Clam:      &FreshClam{},
		Veggies:   []Veggieer{&Garlic{}, &Onion{}, &Mushroom{}, &RedPepper{}},
	},
	}
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
