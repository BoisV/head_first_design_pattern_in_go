package rpg2

type WeaponBehavior interface {
	Attack()
	Display()
}

type ArmorBehavior interface {
}

type SwordBehavior interface {
	Chop()
	Attack()
	Display()
}

type ShieldBehavior interface {
	Resist()
}

type StaffBehavior interface {
	Attack()
	Display()
	Magic()
}
