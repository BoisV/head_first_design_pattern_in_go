package simuduck

import "fmt"

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type DisplayBehavior interface {
	Display()
}

type FlyWithWings struct {
}

func (f *FlyWithWings) Fly() {
	fmt.Println("I'm flying!!")
}

type FlyWithNoWay struct {
}

func (f *FlyWithNoWay) Fly() {
	fmt.Println("I can't fly")
}

type MuteQuack struct {
}

func (m *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct {
}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}
