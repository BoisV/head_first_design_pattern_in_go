package starbuzz_coffee

import (
	"fmt"
	"testing"
)

func TestStarbuzz_Coffee(t *testing.T) {
	milk := NewMilk(NewHouseBlend())
	fmt.Printf("%s $%.2f\n", milk.GetDescription(), milk.Cost())
	milk.Beverage = nil
	fmt.Printf("%s $%.2f\n", milk.GetDescription(), milk.Cost())

}

func TestStarbuzz_Coffee2(t *testing.T) {
	whip := NewWhip(NewMocha(NewDarkRoast()))
	fmt.Printf("%s $%.2f\n", whip.GetDescription(), whip.Cost())

}

func TestStarbuzz_Coffee3(t *testing.T) {
	mocha := NewMocha(NewMocha(NewSoy(NewWhip(nil))))
	fmt.Printf("%s $%.2f\n", mocha.GetDescription(), mocha.Cost())

}
