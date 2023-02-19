package starbuzz_coffee

import (
	"testing"
)

func TestStarbuzz_Coffee(t *testing.T) {
	milk := Milk{CondimentDecorator{beverage: HouseBlend{}, price: 10, description: "milk"}}
	println(milk.Cost(), "元")
	println(milk.GetDescription())
	milk.SetBeverage(nil)
	println(milk.Cost(), "元")
	println(milk.GetDescription())
}
