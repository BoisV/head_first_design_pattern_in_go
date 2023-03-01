package remote_control

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SimpleRemoteControl struct {
	slot Command
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	return &SimpleRemoteControl{}
}

func (s *SimpleRemoteControl) SetSlot(slot Command) {
	s.slot = slot
}

func (s *SimpleRemoteControl) ButtonWasPressed() {
	s.slot.Execute()
}

type RemoteControlSuite struct {
	suite.Suite
}

func TestRemoteControlSuite(t *testing.T) {
	suite.Run(t, new(RemoteControlSuite))
}

func (su *RemoteControlSuite) TestSimpleRemoteControl() {
	light := NewLight()
	garageDoor := NewGarageDoor()
	remoteControl := NewSimpleRemoteControl()
	su.Run("light on", func() {
		remoteControl.SetSlot(NewLightOnCommand(light))
		remoteControl.ButtonWasPressed()
		su.Equal(true, light.Status)

	})
	su.Run("light off", func() {
		remoteControl.SetSlot(NewLightOffCommand(light))
		remoteControl.ButtonWasPressed()
		su.Equal(false, light.Status)
	})

	su.Run("garage door open", func() {
		remoteControl.SetSlot(NewGarageDoorOpenCommand(garageDoor))
		remoteControl.ButtonWasPressed()
		su.Equal(true, garageDoor.doorStatus == UP)
	})
}

func (su *RemoteControlSuite) TestRemoteControl_String() {
	su.Run("string", func() {
		remoteControl := NewRemoteControl()
		su.Equal(
			`
------ Remote Control ------
[slot 0] NoCommand    NoCommand
[slot 1] NoCommand    NoCommand
[slot 2] NoCommand    NoCommand
[slot 3] NoCommand    NoCommand
[slot 4] NoCommand    NoCommand
[slot 5] NoCommand    NoCommand
[slot 6] NoCommand    NoCommand
`, remoteControl.String())
	})

	su.Run("remote control", func() {
		control := NewRemoteControl()
		light := NewLight()
		control.SetCommand(0, NewLightOnCommand(light), NewLightOffCommand(light))
		control.OnButtonWasPressed(0)
		su.Equal(true, light.Status)
	})
}
