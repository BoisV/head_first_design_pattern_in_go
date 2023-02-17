package rpg

import "fmt"

type WeaponBehavior interface {
	UseWeapon() string
}

// Sword 剑
type Sword struct {
	Name string
	ATK  int32
}

func (s *Sword) UseWeapon() string {
	return fmt.Sprintf("挥舞<<%s>>(剑)对敌人造成了%d点伤害", s.Name, s.ATK)
}

// Axe 斧
type Axe struct {
	Name string
	ATK  int32
}

func (a *Axe) UseWeapon() string {
	return fmt.Sprintf("挥舞<<%s>>(斧)对敌人造成了%d点伤害", a.Name, a.ATK)
}

// BowAndArrow 弓箭
type BowAndArrow struct {
	Name string
	ATK  int32
}

func (b *BowAndArrow) UseWeapon() string {
	return fmt.Sprintf("使用<<%s>>(弓箭)对敌人造成了%d点伤害", b.Name, b.ATK)
}

// Knife 小刀
type Knife struct {
	Name string
	ATK  int32
}

func (k *Knife) UseWeapon() string {
	return fmt.Sprintf("挥舞<<%s>>(小刀)对敌人造成了%d点伤害", k.Name, k.ATK)
}
