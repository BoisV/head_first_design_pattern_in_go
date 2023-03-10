package duck

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DuckSuite struct {
	suite.Suite
}

func TestDuckSuite(t *testing.T) {
	suite.Run(t, new(DuckSuite))
}

func (su *DuckSuite) TestTurkeyAdapter() {
	su.Run("adapter test", func() {
		duck := NewMallardDuck()

		wildTurkey := NewWildTurkey()
		turkeyAdapter := NewTurkeyAdapter(wildTurkey)

		fmt.Println("The Turkey says...")
		wildTurkey.Gobble()
		wildTurkey.Fly()

		fmt.Println("\nThe Duck says...")
		duck.Quack()
		duck.Fly()

		fmt.Println("\nThe TurkeyAdapter says...")
		turkeyAdapter.Quack()
		turkeyAdapter.Fly()
	})
}
