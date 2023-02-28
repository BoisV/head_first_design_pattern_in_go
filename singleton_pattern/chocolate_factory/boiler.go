package chocolate_factory

import "sync"

type boiler struct {
	empty  bool
	boiled bool
}

func NewBoiler() *boiler {
	return &boiler{empty: true, boiled: false}
}

func (b *boiler) Fill() {
	if b.IsEmpty() {
		b.empty = false
		b.boiled = false
		// fill the boiler with a milk/chocolate mixture
	}
}

func (b *boiler) Drain() {
	if !b.IsEmpty() && b.IsBoiled() {
		// drain the boiled milk and chocolate
		b.empty = true
	}
}

func (b *boiler) Boil() {
	if !b.IsEmpty() && !b.IsBoiled() {
		// bring the contents to a boil
		b.boiled = true
	}
}

func (b *boiler) IsEmpty() bool {
	return b.empty
}

func (b *boiler) IsBoiled() bool {
	return b.boiled
}

var lock = &sync.Mutex{}
var boilerInstance *boiler

func GetBoilerInstance() *boiler {
	if boilerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if boilerInstance == nil {
			boilerInstance = NewBoiler()
		}
	}
	return boilerInstance
}
