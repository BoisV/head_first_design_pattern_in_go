package rpg

import (
	"testing"
)

func TestRPGGame(t *testing.T) {
	king := King{Character{Name: "国王"}}
	king.SetWeapon(Axe{Name: "血吼", ATK: 30})
	king.fight()

	queen := Queen{Character{Name: "王后"}}
	queen.SetWeapon(BowAndArrow{Name: "轻语之弓", ATK: 15})
	queen.fight()

	troll := Troll{Character{Name: "巨魔"}}
	troll.SetWeapon(Knife{Name: "多兰剑", ATK: 20})
	troll.fight()

	knight := Knight{Character{Name: "骑士"}}
	knight.SetWeapon(Sword{Name: "霜之哀伤", ATK: 25})
	knight.fight()
}
