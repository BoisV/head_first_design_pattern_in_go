package remote_control

import (
	"fmt"
	"reflect"
	"strings"
)

type RemoteControl struct {
	OnCommands [7]Command
	OffCommand [7]Command
}

func (r *RemoteControl) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n------ Remote Control ------\n")
	for i := 0; i < len(r.OnCommands); i++ {
		builder.WriteString(fmt.Sprintf("[slot %d] %s    %s\n", i, reflect.TypeOf(r.OnCommands[i]).Elem().Name(), reflect.TypeOf(r.OnCommands[i]).Elem().Name()))
	}
	return builder.String()
}

type NoCommand struct {
}

func NewNoCommand() *NoCommand {
	return &NoCommand{}
}

func (c *NoCommand) Execute() {

}

func NewRemoteControl() *RemoteControl {
	r := &RemoteControl{
		OnCommands: [7]Command{},
		OffCommand: [7]Command{},
	}
	for i := 0; i < 7; i++ {
		r.OnCommands[i] = NewNoCommand()
		r.OffCommand[i] = NewNoCommand()
	}
	return r
}

func (r *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	r.OnCommands[slot] = onCommand
	r.OffCommand[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPressed(slot int) {
	r.OnCommands[slot].Execute()
}

func (r *RemoteControl) OffButtonWasPressed(slot int) {
	r.OffCommand[slot].Execute()
}
