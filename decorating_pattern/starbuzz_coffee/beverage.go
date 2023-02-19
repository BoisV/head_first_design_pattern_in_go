package starbuzz_coffee

type Beverage interface {
	Cost() int64
	GetDescription() string
}

type DarkRoast struct {
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{}
}

func (d DarkRoast) Cost() int64 {
	return 10
}

func (d DarkRoast) GetDescription() string {
	return "whip"
}

type HouseBlend struct {
}

func (h HouseBlend) Cost() int64 {
	return 19
}

func (h HouseBlend) GetDescription() string {
	return "house blend"
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{}
}

type Espresso struct {
}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e Espresso) Cost() int64 {
	return 30
}

func (e Espresso) GetDescription() string {
	return "espresso"
}

type Decaf struct {
}

func (d Decaf) Cost() int64 {
	return 25
}

func (d Decaf) GetDescription() string {
	return "decaf"
}

type CondimentDecorator struct {
	beverage    Beverage
	price       int64
	description string
}

func (b *CondimentDecorator) GetDescription() string {
	if b.beverage != nil {
		return b.beverage.GetDescription() + " and " + b.description
	}
	return b.description
}

func (b *CondimentDecorator) Cost() int64 {
	if b.beverage == nil {
		return b.price
	}
	return b.beverage.Cost() + b.price
}

func (b *CondimentDecorator) SetBeverage(beverage Beverage) {
	b.beverage = beverage
}

type Mocha struct {
	CondimentDecorator
}

func NewMocha(condimentDecorator CondimentDecorator) *Mocha {
	return &Mocha{CondimentDecorator: condimentDecorator}
}

func (m Mocha) Cost() int64 {
	return m.CondimentDecorator.Cost()
}

func (m Mocha) GetDescription() string {
	return m.CondimentDecorator.GetDescription()
}

type Whip struct {
	CondimentDecorator
}

func NewWhip(condimentDecorator CondimentDecorator) *Whip {
	return &Whip{CondimentDecorator: condimentDecorator}
}

func (w Whip) Cost() int64 {
	return w.CondimentDecorator.Cost()
}

func (w Whip) GetDescription() string {
	return w.CondimentDecorator.GetDescription()
}

type Milk struct {
	CondimentDecorator
}

func (m Milk) Cost() int64 {
	return m.CondimentDecorator.Cost()
}

func (m Milk) GetDescription() string {
	return m.CondimentDecorator.GetDescription()
}

func NewMilk() *Milk {
	return &Milk{}
}

type Soy struct {
	CondimentDecorator
}

func NewSoy(condimentDecorator CondimentDecorator) *Soy {
	return &Soy{CondimentDecorator: condimentDecorator}
}

func (s Soy) Cost() int64 {
	return s.CondimentDecorator.Cost()
}

func (s Soy) GetDescription() string {
	return s.CondimentDecorator.GetDescription()
}
