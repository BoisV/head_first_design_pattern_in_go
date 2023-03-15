package caffeine_beverage

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func NewTea() *Tea {
	return &Tea{}
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}
