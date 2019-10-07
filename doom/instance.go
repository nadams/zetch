package doom

type State string

const (
	Running State = "running"
	Stopped
)

type Instance struct {
	conf    Config
	retries int
	state   State
}

func NewInstance(c Config) *Instance {
	return &Instance{
		conf: c,
	}
}

func (i *Instance) Start() error {
	return nil

}
