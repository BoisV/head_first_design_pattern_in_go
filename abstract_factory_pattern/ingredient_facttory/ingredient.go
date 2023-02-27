package ingredient_facttory

// Dougher this interface defines methods of Dougher
type Dougher interface {
	GetName() string
}

type Dough struct {
	name string
}

func NewDough(name string) *Dough {
	return &Dough{name: name}
}

func (d *Dough) GetName() string {
	return d.name
}

type ThinCrustDough struct {
	Dough
}

func NewThinCrustDough() *ThinCrustDough {
	return &ThinCrustDough{Dough: Dough{name: "Thin Crust Dough"}}
}

// Saucer this interface defines methods of Saucer
type Saucer interface {
	GetName() string
}
type Sauce struct {
	name string
}

func NewSauce(name string) *Sauce {
	return &Sauce{name: name}
}

func (s *Sauce) GetName() string {
	return s.name
}

type MarinaraSauce struct {
	Sauce
}

func NewMarinaraSauce() *MarinaraSauce {
	return &MarinaraSauce{Sauce: Sauce{name: "Marinara Sauce"}}
}

// Cheeser this interface defines methods of Cheeser
type Cheeser interface {
	GetName() string
}
type Cheese struct {
	name string
}

func NewCheese(name string) *Cheese {
	return &Cheese{name: name}
}

func (c *Cheese) GetName() string {
	return c.name
}

type ReggianoCheese struct {
	Cheese
}

func NewReggianoCheese() *ReggianoCheese {
	return &ReggianoCheese{Cheese: Cheese{name: "Regginao Cheese"}}
}

// Pepperonier this interface defines methods of Pepperonier
type Pepperonier interface {
	GetName() string
}
type Pepperoni struct {
	name string
}

func NewPepperoni(name string) *Pepperoni {
	return &Pepperoni{name: name}
}

func (p *Pepperoni) GetName() string {
	return p.name
}

type SlicePepperoni struct {
	Pepperoni
}

func NewSlicePepperoni() *SlicePepperoni {
	return &SlicePepperoni{Pepperoni: Pepperoni{name: "Slice Pepperoni"}}
}

// Clamer this interface defines methods of Clam
type Clamer interface {
	GetName() string
}

type Clam struct {
	name string
}

func NewClam(name string) *Clam {
	return &Clam{name: name}
}

func (c *Clam) GetName() string {
	return c.name
}

type FreshClam struct {
	Clam
}

func NewFreshClam() *FreshClam {
	return &FreshClam{Clam: Clam{name: "Fresh Clam"}}
}

// Veggieer this interface defines methods of Veggie
type Veggieer interface {
	GetName() string
}

type Veggie struct {
	name string
}

func NewVeggie(name string) *Veggie {
	return &Veggie{name: name}
}

func (v *Veggie) GetName() string {
	return v.name
}

type Onion struct {
	Veggie
}

func NewOnion() *Onion {
	return &Onion{Veggie: Veggie{name: "Onion"}}
}

type Mushroom struct {
	Veggie
}

func NewMushroom() *Mushroom {
	return &Mushroom{Veggie: Veggie{name: "Mushroom"}}
}

type RedPepper struct {
	Veggie
}

func NewRedPepper() *RedPepper {
	return &RedPepper{Veggie: Veggie{name: "Red Pepper"}}
}

type Garlic struct {
	Veggie
}

func NewGarlic() *Garlic {
	return &Garlic{Veggie: Veggie{name: "Garlic"}}
}
