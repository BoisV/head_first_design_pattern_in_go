package rpg2

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Character struct {
	Weapon WeaponBehavior
}

type Rpg2Suite struct {
	suite.Suite
}

func TestRpg2Suite(t *testing.T) {
	suite.Run(t, new(Rpg2Suite))
}

func (su *Rpg2Suite) TestRPGGame() {
	su.Run("剑", func() {
		c := Character{Weapon: &Sword{Name: "霜之哀伤", ATK: 999}}
		c.Weapon.Display()
		c.Weapon.Attack()

		c.Weapon = &Sword{Name: "火之高兴", ATK: 888}
		c.Weapon.Display()
		c.Weapon.Attack()

		c.Weapon = &Staff{Name: "鸡腿杖", MagicType: "火焰", ATK: 888}
		c.Weapon.Display()
		c.Weapon.Attack()

		c.Weapon = &Staff{Name: "远古法杖", MagicType: "砂石", ATK: 888}
		c.Weapon.Display()
		c.Weapon.Attack()
	})

}
