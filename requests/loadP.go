package requests

import (
	"sysops/globals"
	"sysops/monitor"
)

//FreeMem request to free mem
type LoadP struct {
	Type  string
	Input []string
	Size  int
	PID   int
	Logs  []*monitor.PageLog
}

//NewFreeMemReq n
func NewLoadP(input []string, PID int, size int) *LoadP {

	return &LoadP{
		Type:  globals.LoadP,
		Input: input,
		PID:   PID,
		Size:  size,
		Logs:  []*monitor.PageLog{},
	}
}

//GenerateOutput p
func (f *LoadP) GenerateOutput() {

}

//Args a
func (f *LoadP) Args() []int {
	args := []int{f.PID, f.Size}
	return args
}

//Type t
func (f *LoadP) GetType() string {
	return f.Type
}
