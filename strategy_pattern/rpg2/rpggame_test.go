package rpg2

import "testing"

type Character struct {
	Weapon WeaponBehavior
}

func TestRPGGame(t *testing.T) {
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
}
