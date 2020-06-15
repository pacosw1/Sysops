package requests

import (
	"sysops/globals"
	"sysops/monitor"
)

//FreeMem request to free mem
type AccessReq struct {
	Type  string
	Input []string
	PID   int
	VAddr int
	Mod   int
	Logs  []*monitor.PageLog
}

//NewFreeMemReq n
func NewAccessReq(input []string, PID, vAddr, mod int) *AccessReq {

	return &AccessReq{
		Type:  globals.Access,
		Input: input,
		PID:   PID,
		VAddr: vAddr,
		Mod:   mod,
		Logs:  []*monitor.PageLog{},
	}
}

//GenerateOutput p
func (f *AccessReq) GenerateOutput() {

}

//Args a
func (f *AccessReq) Args() []int {
	args := []int{f.PID, f.VAddr, f.Mod}
	return args
}

//Type t
func (f *AccessReq) GetType() string {
	return f.Type
}

//Logs t
func (f *AccessReq) GetLogs() []*monitor.PageLog {
	return f.Logs
}
