package home_sweet_home_theater

import "fmt"

type Screen struct {
}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) up() {
	fmt.Println("Theater Screen going up")
}

func (s *Screen) down() {
	fmt.Println("Theater Screen going down")
}

type PopcornPopper struct {
}

func NewPopcornPopper() *PopcornPopper {
	return &PopcornPopper{}
}

func (p *PopcornPopper) pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *PopcornPopper) on() {
	fmt.Println("Popcorn Popper on")
}

func (p *PopcornPopper) off() {
	fmt.Println("Popcorn Popper off")
}

type TheaterLights struct {
}

func NewTheaterLights() *TheaterLights {
	return &TheaterLights{}
}

func (t *TheaterLights) dim(v int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", v)
}

func (t *TheaterLights) on() {
	fmt.Println("Theater Ceiling Lights on")
}

func (t *TheaterLights) off() {
	fmt.Println("Theater Ceiling Lights off")
}
