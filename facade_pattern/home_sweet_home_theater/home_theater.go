package home_sweet_home_theater

import "fmt"

type HomeTheaterFacade struct {
	Amp       IAmplifier
	Tuner     ITuner
	Player    IPlayer
	Projector IProjector
	Lights    *TheaterLights
	Screen    *Screen
	Popper    *PopcornPopper
}

func NewHomeTheaterFacade(amplifier IAmplifier,
	tuner ITuner,
	streamingPlayer IPlayer,
	projector IProjector,
	theaterLights *TheaterLights,
	screen *Screen,
	popcornPopper *PopcornPopper) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		Amp:       amplifier,
		Tuner:     tuner,
		Player:    streamingPlayer,
		Projector: projector,
		Lights:    theaterLights,
		Screen:    screen,
		Popper:    popcornPopper}
}

func (t *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	t.Popper.on()
	t.Popper.pop()
	t.Lights.dim(10)
	t.Screen.down()
	t.Projector.on()
	t.Projector.wideScreenMode()
	t.Amp.on()
	t.Amp.setStreamingPlayer(t.Player)
	t.Amp.setVolume(5)
	t.Amp.setSurroundSound()
	t.Player.on()
	t.Player.play(movie)
}

func (t *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	t.Popper.off()
	t.Lights.on()
	t.Screen.up()
	t.Projector.off()
	t.Amp.off()
	t.Player.stop()
	t.Player.off()
}
