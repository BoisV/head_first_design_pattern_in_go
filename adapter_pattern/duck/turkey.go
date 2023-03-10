package duck

import "fmt"

type ITurkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
}

func NewWildTurkey() *WildTurkey {
	return &WildTurkey{}
}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}

type TurkeyAdapter struct {
	turkey ITurkey
}

func NewTurkeyAdapter(iTurkey ITurkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: iTurkey}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}
