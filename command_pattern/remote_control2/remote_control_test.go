package remote_control

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RemoteControlSuite struct {
	suite.Suite
}

func TestRemoteControlSuite(t *testing.T) {
	suite.Run(t, new(RemoteControlSuite))
}

func (su *RemoteControlSuite) TestRemoteControl_String() {

	su.Run("remote control", func() {
		control := NewRemoteControl()
		light := NewLight()
		control.SetCommand(0, func() {
			light.On()
		}, func() {
			light.Off()
		})
		control.OnButtonWasPressed(0)
		su.Equal(true, light.Status)
	})

	su.Run("undo", func() {
		control := NewRemoteControl()
		light := NewLight()
		control.SetCommand(0, func() { light.On() }, func() { light.Off() })
		control.OnButtonWasPressed(0)
		su.Equal(true, light.Status)
		control.Undo()
		su.Equal(false, light.Status)
	})
}
