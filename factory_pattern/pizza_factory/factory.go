package pizza_factory

type PizzaFactorier interface {
	CreatePizza(pizzaType string) Pizzaer
}

type PizzaFactory struct {
}

func (p *PizzaFactory) CreatePizza(pizzaType string) Pizzaer {
	var pizza Pizzaer = nil
	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "clam":
		pizza = NewClamPizza()
	case "veggie":
		pizza = NewVeggiePizza()

	}
	return pizza
}
