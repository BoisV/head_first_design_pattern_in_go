package remote_control

import "fmt"

const (
	DOWN = iota
	UP
	STOP
)

type Light struct {
	Status bool
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) On() {
	l.Status = true
	fmt.Println("light on")
}

func (l *Light) Off() {
	l.Status = false
	fmt.Println("light off")
}

type GarageDoor struct {
	doorStatus int8 // 0:down，1:up, 2:stop
	light      Light
}

func NewGarageDoor() *GarageDoor {
	return &GarageDoor{}
}
func (d *GarageDoor) Down() {
	d.doorStatus = DOWN
}

func (d *GarageDoor) Up() {
	d.doorStatus = UP
}

func (d *GarageDoor) Stop() {
	d.doorStatus = STOP
}

func (d *GarageDoor) LightOn() {
	d.light.On()
}

func (d *GarageDoor) LightOff() {
	d.light.Off()
}

type Stereo struct {
	volume int
}

func (s *Stereo) On() {
	fmt.Println("stereo on")
}

func (s *Stereo) off() {
	fmt.Println("stereo off")
}

func (s *Stereo) setCD(name string) {
	fmt.Printf("stereo play CD <%s>\n", name)
}

func (s *Stereo) setDvd(name string) {
	fmt.Printf("stereo play DVD <%s>\n", name)
}

func (s *Stereo) setRadio() {
	fmt.Println("stereo plays radio")
}

func (s *Stereo) setVolume(v int) {
	s.volume = v
}
