package ingredient_facttory

type PizzaStore struct {
	PizzaIngredientFactory PizzaIngredientFactorier
}

func (s *PizzaStore) CreatePizza(pizzaType string) Pizzaer {
	panic("implement this...")
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

// NYPizzaStore 纽约披萨店
type NYPizzaStore struct {
	PizzaStore
}

func NewNYPizzaStore(factory PizzaIngredientFactorier) *NYPizzaStore {
	return &NYPizzaStore{PizzaStore: PizzaStore{PizzaIngredientFactory: factory}}
}

func (s *NYPizzaStore) CreatePizza(pizzaType string) Pizzaer {
	switch pizzaType {
	case "cheese":
		pizza := NewCheesePizza(s.PizzaIngredientFactory)
		pizza.Name = "New York style Cheese Pizza"
		return pizza
	case "veggie":
		pizza := NewVeggiePizza(s.PizzaIngredientFactory)
		pizza.Name = "New York style Veggie Pizza"
		return pizza
	case "clam":
		pizza := NewClamPizza(s.PizzaIngredientFactory)
		pizza.Name = "New York style Clam Pizza"
		return pizza
	case "pepperoni":
		pizza := NewPepperoniPizza(s.PizzaIngredientFactory)
		pizza.Name = "New York style Pepperoni Pizza"
		return pizza
	}
	return nil
}

func (s *NYPizzaStore) OrderPizza(pizzaType string) Pizzaer {
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

// ChicagoPizzaStore 芝加哥披萨店
type ChicagoPizzaStore struct {
	PizzaStore
}

func NewChicagoPizzaStore(factory PizzaIngredientFactorier) *ChicagoPizzaStore {
	return &ChicagoPizzaStore{PizzaStore{PizzaIngredientFactory: factory}}
}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) Pizzaer {
	switch pizzaType {
	case "cheese":
		pizza := NewCheesePizza(s.PizzaIngredientFactory)
		pizza.Name = "Chicago style Cheese Pizza"
		return pizza

	case "pepperoni":
		pizza := NewPepperoniPizza(s.PizzaIngredientFactory)
		pizza.Name = "Chicago style Pepperoni Pizza"
		return pizza

	case "clam":
		pizza := NewClamPizza(s.PizzaIngredientFactory)
		pizza.Name = "Chicago style Clam Pizza"
		return pizza

	case "veggie":
		pizza := NewVeggiePizza(s.PizzaIngredientFactory)
		pizza.Name = "Chicago style Veggie Pizza"
		return pizza
	}
	return nil
}

func (s *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizzaer {
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
