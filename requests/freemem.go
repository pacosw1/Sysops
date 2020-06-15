package requests

import (
	"sysops/globals"
	"sysops/monitor"
)

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
		Logs:  []*monitor.PageLog{},
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
func (f *FreeMem) GetType() string {
	return f.Type
}
