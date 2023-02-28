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
var boilerInstanceDCL *boiler

// GetBoilerInstanceByDoubelCheckedLock 双重校验锁
func GetBoilerInstanceByDoubelCheckedLock() *boiler {
	if boilerInstanceDCL == nil {
		lock.Lock()
		defer lock.Unlock()
		if boilerInstanceDCL == nil {
			boilerInstanceDCL = NewBoiler()
		}
	}
	return boilerInstanceDCL
}

var boilerInstanceEagerly = new(boiler)

// GetBoilerInstanceEagerly 饿汉式
func GetBoilerInstanceEagerly() *boiler {
	return boilerInstanceEagerly
}

var lock2 = &sync.Mutex{}
var boilerInstanceLazily *boiler

// GetBoilerInstanceLazily 懒汉式
func GetBoilerInstanceLazily() *boiler {
	lock2.Lock()
	defer lock2.Unlock()
	if boilerInstanceLazily == nil {
		boilerInstanceLazily = new(boiler)
	}
	return boilerInstanceLazily
}
