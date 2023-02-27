package ingredient_facttory

// PizzaIngredientFactorier 定义披萨原料工厂的方法
type PizzaIngredientFactorier interface {
	CreateDough() Dougher
	CreateSauce() Saucer
	CreateCheese() Cheeser
	CreateVeggie() []Veggieer
	CreatePepperoni() Pepperonier
	CreateClam() Clamer
}

type PizzaIngredientFactory struct {
}

func (p *PizzaIngredientFactory) CreateDough() Dougher {
	return NewThinCrustDough()
}

func (p *PizzaIngredientFactory) CreateSauce() Saucer {
	return NewMarinaraSauce()
}

func (p *PizzaIngredientFactory) CreateCheese() Cheeser {
	return NewReggianoCheese()
}

func (p *PizzaIngredientFactory) CreateVeggie() []Veggieer {
	return []Veggieer{NewMushroom(), NewGarlic(), NewOnion(), NewRedPepper()}
}

func (p *PizzaIngredientFactory) CreatePepperoni() Pepperonier {
	return NewSlicePepperoni()
}

func (p *PizzaIngredientFactory) CreateClam() Clamer {
	return NewFreshClam()
}
