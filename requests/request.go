package requests

//Request command input request interface
type Request interface {
	GenerateOutput(output string)
	GetType() string
	AddLog(*PageLog)

	// GetLogs() []*PageLog
	Args() []int
}
