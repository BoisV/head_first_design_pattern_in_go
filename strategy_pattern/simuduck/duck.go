package simuduck

import "fmt"

type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.QuackBehavior.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) Display() {
	fmt.Println("I'm a duck")
}
