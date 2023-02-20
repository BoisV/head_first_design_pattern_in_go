package rpg2

import "fmt"

// Sword 武器-剑
type Sword struct {
	Name string
	ATK  int32
}

func (s *Sword) Attack() {
	s.Chop()
}

func (s *Sword) Display() {
	fmt.Println(fmt.Sprintf("武器名称是 \"%s\"， 种类是剑", s.Name))
}

func (s *Sword) Chop() {
	fmt.Println(fmt.Sprintf("敌人被砍伤 掉 %d 血 ", s.ATK))
}

// Staff 法杖
type Staff struct {
	Name      string
	ATK       int32
	MagicType string
}

func (s *Staff) Attack() {
	s.Magic()
}

func (s *Staff) Display() {
	fmt.Println(fmt.Sprintf("武器名称是 \"%s\"， 种类是法杖", s.Name))
}

func (s *Staff) Magic() {
	fmt.Println(fmt.Sprintf("敌人遭到%s魔法 掉 %d 血 ", s.MagicType, s.ATK))
}
