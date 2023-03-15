package home_sweet_home_theater

import "fmt"

type ITuner interface {
	setAm()
	setFm()
	setFrequency()
}

type Tuner struct {
	Amplifier IAmplifier
}

func (t *Tuner) on() {
	fmt.Println("Tuner on")
}

func (t *Tuner) off() {
	fmt.Println("Tuner off")
}

func (t *Tuner) setAm() {
	//TODO implement me
	panic("implement me")
}

func (t *Tuner) setFm() {
	//TODO implement me
	panic("implement me")
}

func (t *Tuner) setFrequency() {
	//TODO implement me
	panic("implement me")
}
