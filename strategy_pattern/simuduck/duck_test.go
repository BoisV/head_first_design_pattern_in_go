package simuduck

import (
	"fmt"
	"testing"
)

type ModelDuck struct {
	Duck
}

func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

type FlyRocketPowered struct {
}

func (f *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}

func TestDuck(t *testing.T) {
	var modelDuck ModelDuck = ModelDuck{
		Duck: Duck{FlyBehavior: &FlyWithWings{}, QuackBehavior: &MuteQuack{}},
	}
	modelDuck.FlyBehavior = &FlyRocketPowered{}

	modelDuck.Display()
	modelDuck.PerformFly()
	modelDuck.PerformQuack()
}
