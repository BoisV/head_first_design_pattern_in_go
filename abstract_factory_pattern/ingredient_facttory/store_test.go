package ingredient_facttory

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PizzaStoreSuite struct {
	suite.Suite
	PizzaIngredientFactory PizzaIngredientFactory
}

func TestPizzaStoreSuite(t *testing.T) {
	suite.Run(t, new(PizzaStoreSuite))
}

func (su *PizzaStoreSuite) SetupTest() {
	su.PizzaIngredientFactory = PizzaIngredientFactory{}
}

func (su *PizzaStoreSuite) TestOrderPizza() {
	su.Run("从纽约店订购披萨", func() {
		nyPizzaStore := NewNYPizzaStore(&su.PizzaIngredientFactory)
		pizza := nyPizzaStore.OrderPizza("cheese")
		su.Equal("New York style Cheese Pizza", pizza.GetName())
	})

	su.Run("从纽约店订购披萨失败", func() {
		nyPizzaStore := NewNYPizzaStore(&su.PizzaIngredientFactory)
		pizza := nyPizzaStore.OrderPizza("bean")
		su.Nil(pizza)
	})

	su.Run("从芝加哥店订购披萨", func() {
		chicagoPizzaStore := NewChicagoPizzaStore(&su.PizzaIngredientFactory)
		pizza := chicagoPizzaStore.OrderPizza("cheese")
		su.Equal("Chicago style Cheese Pizza", pizza.GetName())
	})

	su.Run("从芝加哥店订购披萨失败", func() {
		chicagoPizzaStore := NewChicagoPizzaStore(&su.PizzaIngredientFactory)
		pizza := chicagoPizzaStore.OrderPizza("bean")
		su.Nil(pizza)
	})
}
