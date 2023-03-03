package remote_control

type Command func()
type UndoCommand func()

type RemoteControl struct {
	OnCommands  [7]Command
	OffCommand  [7]Command
	UndoCommand UndoCommand
}

func NewRemoteControl() *RemoteControl {
	r := &RemoteControl{
		OnCommands: [7]Command{},
		OffCommand: [7]Command{},
	}
	for i := 0; i < 7; i++ {
		r.OnCommands[i] = func() {}
		r.OffCommand[i] = func() {}
	}
	return r
}

func (r *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	r.OnCommands[slot] = onCommand
	r.OffCommand[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPressed(slot int) {
	r.OnCommands[slot]()
	r.UndoCommand = UndoCommand(r.OffCommand[slot])
}

func (r *RemoteControl) OffButtonWasPressed(slot int) {
	r.OffCommand[slot]()
	r.UndoCommand = UndoCommand(r.OnCommands[slot])
}

func (r *RemoteControl) Undo() {
	r.UndoCommand()
}
