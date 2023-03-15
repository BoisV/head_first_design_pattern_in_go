package home_sweet_home_theater

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type HomeTheaterSuite struct {
	suite.Suite
	amp       IAmplifier
	player    IPlayer
	lights    *TheaterLights
	projector IProjector
	screen    *Screen
	popper    *PopcornPopper
}

func TestHomeTheater(t *testing.T) {
	suite.Run(t, new(HomeTheaterSuite))
}
func (su *HomeTheaterSuite) SetupTest() {
	su.amp = NewAmplifier(nil, nil)
	su.player = NewStreamingPlayer(su.amp)
	su.amp.setStreamingPlayer(su.player)
	su.lights = NewTheaterLights()
	su.projector = NewProjector(su.player)
	su.screen = NewScreen()
	su.popper = NewPopcornPopper()
}

func (su *HomeTheaterSuite) TestHomeTheaterFacade() {

	su.Run("HomeTheaterFacade", func() {
		facade := NewHomeTheaterFacade(su.amp, nil, su.player, su.projector, su.lights, su.screen, su.popper)
		facade.WatchMovie("Raider of the Lost Ark")
		facade.EndMovie()
	})
}
