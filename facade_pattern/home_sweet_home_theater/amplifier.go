package home_sweet_home_theater

import "fmt"

type IAmplifier interface {
	setStereoSound()
	setSurroundSound()
	setVolume(int)
	on()
	off()
	setStreamingPlayer(player IPlayer)
}

type Amplifier struct {
	Tuner  ITuner
	Player IPlayer
}

func NewAmplifier(tuner ITuner, player IPlayer) *Amplifier {
	return &Amplifier{Tuner: tuner, Player: player}
}

func (a *Amplifier) on() {
	fmt.Println("Amplifier on")
}

func (a *Amplifier) off() {
	fmt.Println("Amplifier off")

}

func (a *Amplifier) setStereoSound() {
	fmt.Println("Amplifier stereo sound on")
}

func (a *Amplifier) setSurroundSound() {
	fmt.Println("Amplifier surround sound on (5 speakers, 1 subwoofer")

}

func (a *Amplifier) setVolume(v int) {
	fmt.Println("Amplifier setting volume to", v)
}
func (a *Amplifier) setStreamingPlayer(player IPlayer) {
	a.Player = player
	fmt.Println("Amplifier setting Streaming Player to Player")
}
