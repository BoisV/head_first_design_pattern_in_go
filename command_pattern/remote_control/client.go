package remote_control

type Client interface {
	CreateCommand() Command
}

type Invoker interface {
}

type MyClient struct {
}

func (m *MyClient) CreateCommand() Command {
	//TODO implement me
	panic("implement me")
}
