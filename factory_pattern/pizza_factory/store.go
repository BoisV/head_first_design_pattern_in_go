package pizza_factory

type PizzaStore struct {
	PizzaFactory PizzaFactory
}

func NewStore(pizzaFactory PizzaFactory) *PizzaStore {
	return &PizzaStore{PizzaFactory: pizzaFactory}
}

func (s *PizzaStore) OrderPizza(pizzaType string) Pizzaer {
	pizza := s.CreatePizza(pizzaType)
	if pizza == nil {
		return pizza
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (s *PizzaStore) CreatePizza(pizzaType string) Pizzaer {
	return s.PizzaFactory.CreatePizza(pizzaType)
}

type NYPizzaStore struct {
	PizzaStore
}

func NewNYPizzaStore(factory PizzaFactory) *NYPizzaStore {
	return &NYPizzaStore{PizzaStore{PizzaFactory: factory}}
}

func (s *NYPizzaStore) CreatePizza(pizzaType string) Pizzaer {
	var pizza Pizzaer = nil
	switch pizzaType {
	case "cheese":
		pizza = NewNYStyleCheesePizza()
	case "pepperoni":
		pizza = NewNYStylePepperoniPizza()
	case "clam":
		pizza = NewNYStyleClamPizza()
	case "veggie":
		pizza = NewNYStyleVeggiePizza()

	}
	return pizza
}

func (s *NYPizzaStore) OrderPizza(pizzaType string) Pizzaer {
	pizza := s.CreatePizza(pizzaType)
	if pizza == nil {
		return pizza
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	return pizza
}

type ChicagoPizzaStore struct {
	PizzaStore
}

func NewChicagoPizzaStore(factory PizzaFactory) *ChicagoPizzaStore {
	return &ChicagoPizzaStore{PizzaStore{PizzaFactory: factory}}
}

func (s *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizzaer {
	pizza := s.CreatePizza(pizzaType)
	if pizza == nil {
		return pizza
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	return pizza
}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) Pizzaer {
	var pizza Pizzaer = nil
	switch pizzaType {
	case "cheese":
		pizza = NewChicagoStyleCheesePizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "clam":
		pizza = NewClamPizza()
	case "veggie":
		pizza = NewVeggiePizza()

	}
	return pizza
}
