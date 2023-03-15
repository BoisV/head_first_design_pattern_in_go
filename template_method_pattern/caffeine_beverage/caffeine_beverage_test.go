package caffeine_beverage

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CaffeineBeverageSuite struct {
	suite.Suite
}

func TestCaffeineBeverage(t *testing.T) {
	suite.Run(t, new(CaffeineBeverageSuite))
}

func (su *CaffeineBeverageSuite) TestT() {
	su.Run("tea", func() {
		tea := NewCaffeineBeverage(NewTea())
		tea.PrepareRecipe()
	})

	su.Run("coffee", func() {
		coffee := NewCaffeineBeverage(NewCoffee())
		coffee.PrepareRecipe()
	})
}
