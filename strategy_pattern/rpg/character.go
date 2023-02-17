package rpg

import "fmt"

type Character struct {
	Name   string
	weapon WeaponBehavior
}

func (c *Character) fight() {
	fmt.Println(c.Name, c.weapon.UseWeapon())
}

// Queen 王后
type Queen struct {
	Character
}

func (q *Queen) SetWeapon(a BowAndArrow) {
	q.weapon = &a
}

// King 国王
type King struct {
	Character
}

func (k *King) SetWeapon(a Axe) {
	k.weapon = &a
}

// Troll 巨魔
type Troll struct {
	Character
}

func (t *Troll) SetWeapon(knife Knife) {
	t.weapon = &knife
}

// Knight 骑士
type Knight struct {
	Character
}

func (k *Knight) SetWeapon(sword Sword) {
	k.weapon = &sword
}
