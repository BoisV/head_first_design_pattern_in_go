package pizza_store

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PizzaStoreSuite struct {
	suite.Suite
	PizzaFactory PizzaFactory
}

func TestPizzaStoreSuite(t *testing.T) {
	suite.Run(t, new(PizzaStoreSuite))
}

func (su *PizzaStoreSuite) SetupTest() {
	su.PizzaFactory = PizzaFactory{}
}

func (su *PizzaStoreSuite) TestStoreOrderPizza() {
	su.Run("购买蔬菜披萨", func() {
		store := PizzaStore{PizzaFactory: su.PizzaFactory}
		viggiePizza := store.OrderPizza("veggie")
		su.Equal(viggiePizza, NewVeggiePizza())
	})

	su.Run("无法购买豌豆披萨", func() {
		store := PizzaStore{PizzaFactory: su.PizzaFactory}
		beanPizza := store.OrderPizza("bean")
		su.Nil(beanPizza)
	})
}

func (su *PizzaStoreSuite) TestNYPizzaStore_OrderPizza() {
	su.Run("购买纽约风味起司披萨", func() {
		store := NewNYPizzaStore(su.PizzaFactory)
		cheesePizza := store.OrderPizza("cheese")
		su.Equal(cheesePizza, NewNYStyleCheesePizza())
	})

	su.Run("购买芝加哥风味起司披萨", func() {
		store := NewChicagoPizzaStore(su.PizzaFactory)
		cheesePizza := store.OrderPizza("cheese")
		su.Equal(cheesePizza, NewChicagoStyleCheesePizza())
	})

}
