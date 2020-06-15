package monitor
import (
	"sysops/globals"
	"sysops/monitor"
)

//Request command input request interface
type Request interface {
	GenerateOutput()
}

//NewRequest Initializes a new Request




//FreeMem request to free mem
type FreeMem struct {
	Type  string
	Input []string
	PID   int
	Logs  []*monitor.PageLog
}

//NewFreeMemReq n
func NewFreeMemReq(input []string, PID int) *FreeMem {

	return &FreeMem{
		Type:  globals.FreeP,
		Input: input,
		PID:   PID,
		Logs:  []*PageLog{},
	}
}


//GenerateOutput p
func (f *FreeMem) GenerateOutput() {

}

//Args a
func (f *FreeMem) Args() []int {
	args := []int{f.PID}
	return args
}

//Type t
func (f *FreeMem) Type() string {
	return f.Type
}


//GenerateOutput p
func (f *FreeMem) GenerateOutput() {

}

//Args a
func (f *FreeMem) Args() []int {
	args := []int{f.PID}
	return args
}

//Type t
func (f *FreeMem) Type() string {
	return f.Type
}








//NewFreeMemReq n
func NewFreeMemReq(input []string, PID int) *FreeMem {

	return &FreeMem{
		Type:  globals.FreeP,
		Input: input,
		PID:   PID,
		Logs:  []*PageLog{},
	}
}

//GenerateOutput p
func (f *FreeMem) GenerateOutput() {

}

//Args a
func (f *FreeMem) Args() []int {
	args := []int{f.PID}
	return args
}

//Type t
func (f *FreeMem) Type() string {
	return f.Type
}





//NewAccessReq n
func NewAccessReq(input []string, PID int) *FreeMem {

	return &FreeMem{
		Type:  globals.FreeP,
		Input: input,
		PID:   PID,
		Logs:  []*PageLog{},
	}
}

//GenerateOutput p
func (f *FreeMem) GenerateOutput() {

}

//Args a
func (f *FreeMem) Args() []int {
	args := []int{f.PID}
	return args
}

//Type t
func (f *FreeMem) Type() string {
	return f.Type
}

