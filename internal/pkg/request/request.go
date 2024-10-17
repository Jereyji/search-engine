package request

type Request struct {
	command string
	flags   map[string]any
}

func NewRequest(command string) *Request {
	return &Request{
		command: command,
		flags:   make(map[string]any),
	}
}

func (r *Request) GetCommand() string {
	return r.command
}

func (r *Request) GetValue(flag string) any {
	return r.flags[flag]
}

func (r *Request) SetCommand(command string) {
	r.command = command
}

func (r *Request) SetFlag(flag string, value any) {
	r.flags[flag] = value
}
