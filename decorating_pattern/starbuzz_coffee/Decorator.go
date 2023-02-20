package starbuzz_coffee

type CondimentDecorator struct {
	Beverage    Beverageer
	price       float64
	description string
}

func (b *CondimentDecorator) GetDescription() string {
	if b.Beverage != nil {
		return b.Beverage.GetDescription() + ", " + b.description
	}
	return b.description
}

func (b *CondimentDecorator) Cost() float64 {
	if b.Beverage == nil {
		return b.price
	}
	return b.Beverage.Cost() + b.price
}

type Mocha struct {
	CondimentDecorator
}

func NewMocha(beverage Beverageer) *Mocha {
	return &Mocha{CondimentDecorator: CondimentDecorator{Beverage: beverage, price: 0.2, description: "mocha"}}
}

type Whip struct {
	CondimentDecorator
}

func NewWhip(beverage Beverageer) *Whip {
	return &Whip{CondimentDecorator: CondimentDecorator{Beverage: beverage, price: 0.1, description: "whip"}}
}

type Milk struct {
	CondimentDecorator
}

func NewMilk(beverage Beverageer) *Milk {
	return &Milk{CondimentDecorator: CondimentDecorator{Beverage: beverage, price: 0.1, description: "milk"}}
}

type Soy struct {
	CondimentDecorator
}

func NewSoy(beverage Beverageer) *Soy {
	return &Soy{CondimentDecorator: CondimentDecorator{Beverage: beverage, price: 0.15, description: "soy"}}
}
