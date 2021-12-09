package guitar

type turning struct{}

func NewTurning() Turning {
	return turning{}
}

func (turning) Threads() []Thread {
	return nil
}
