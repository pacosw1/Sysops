package requests

//Request command input request interface
type Request interface {
	GenerateOutput()
	GetType() string
	Args() []int
}
