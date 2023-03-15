package home_sweet_home_theater

import "fmt"

type IProjector interface {
	tvMode()
	wideScreenMode()
	on()
	off()
}

type Projector struct {
	Player IPlayer
}

func NewProjector(player IPlayer) *Projector {
	return &Projector{Player: player}
}

func (p *Projector) on() {
	fmt.Println("Projector on")
}

func (p *Projector) off() {
	fmt.Println("Projector off")
}

func (p *Projector) tvMode() {
	fmt.Println("Projector is tv mode (4x3 aspect ratio)")
}

func (p *Projector) wideScreenMode() {
	fmt.Println("Projector is widescreen mode (16x9 aspect ratio)")
}
