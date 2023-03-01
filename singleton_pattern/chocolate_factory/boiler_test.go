package chocolate_factory

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ChocolateFactorySuite struct {
	suite.Suite
}

func TestChocolateFactorySuite(t *testing.T) {
	suite.Run(t, new(ChocolateFactorySuite))
}

func (su *ChocolateFactorySuite) TestBoiler() {
	su.Run("双重校验锁", func() {
		instance := GetBoilerInstanceByDoubelCheckedLock()
		instance2 := GetBoilerInstanceByDoubelCheckedLock()
		su.Equal(&instance, &instance2)
	})
	su.Run("饿汉式", func() {
		instance := GetBoilerInstanceEagerly()
		instance2 := GetBoilerInstanceEagerly()
		su.Equal(&instance, &instance2)
	})

	su.Run("懒汉式", func() {
		instance := GetBoilerInstanceLazily()
		instance2 := GetBoilerInstanceLazily()
		su.Equal(&instance, &instance2)
	})
}
