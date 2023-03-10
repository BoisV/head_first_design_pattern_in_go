package duck

import "fmt"

type IDuck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{}
}

func (m *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (m *MallardDuck) Fly() {
	fmt.Println("I'm flying")
}
