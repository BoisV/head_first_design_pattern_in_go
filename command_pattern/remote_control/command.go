package remote_control

type Command interface {
	Execute()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type GarageDoorOpenCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor *GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{garageDoor: garageDoor}
}

func (g *GarageDoorOpenCommand) Execute() {
	g.garageDoor.Up()
}

type StereoOnWithCDCommand struct {
	Stereo Stereo
}

func (c *StereoOnWithCDCommand) Execute() {
	c.Stereo.On()
	c.Stereo.setCD("Wall")
	c.Stereo.setVolume(11)
}
