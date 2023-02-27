package rpg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RPGSuite struct {
	suite.Suite
}

func TestRPGSuite(t *testing.T) {
	suite.Run(t, new(RPGSuite))
}

func (su *RPGSuite) TestRPGGame() {
	su.Run("国王", func() {
		king := King{Character{Name: "国王"}}
		king.SetWeapon(Axe{Name: "血吼", ATK: 30})
		king.fight()
		su.Equal("挥舞<<血吼>>(斧)对敌人造成了30点伤害", king.weapon.UseWeapon())
	})

	su.Run("王后", func() {
		queen := Queen{Character{Name: "王后"}}
		queen.SetWeapon(BowAndArrow{Name: "轻语之弓", ATK: 15})
		queen.fight()
		su.Equal("使用<<轻语之弓>>(弓箭)对敌人造成了15点伤害", queen.weapon.UseWeapon())
	})

	su.Run("巨魔", func() {
		troll := Troll{Character{Name: "巨魔"}}
		troll.SetWeapon(Knife{Name: "多兰剑", ATK: 20})
		troll.fight()
		su.Equal("挥舞<<多兰剑>>(小刀)对敌人造成了20点伤害", troll.weapon.UseWeapon())
	})

	su.Run("骑士", func() {
		knight := Knight{Character{Name: "骑士"}}
		knight.SetWeapon(Sword{Name: "霜之哀伤", ATK: 25})
		knight.fight()
		su.Equal("挥舞<<霜之哀伤>>(剑)对敌人造成了25点伤害", knight.weapon.UseWeapon())
	})

}
