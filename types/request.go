package types

// import (
// 	"sysops/globals"
// 	"sysops/monitor"
// )

// //Request command input request interface
// type Request interface {
// 	GenerateOutput()
// }

// //NewRequest Initializes a new Request

// //FreeMem request to free mem
// type FreeMem struct {
// 	Type  string
// 	Input []string
// 	PID   int
// 	Logs  []*monitor.PageLog
// }

// //GenerateOutput p
// func (f *FreeMem) GenerateOutput() {

// }

// //NewFreeMemReq n
// func NewFreeMemReq(input []string, PID int) *FreeMem {

// 	return &FreeMem{
// 		Type:  globals.FreeP,
// 		Input: input,
// 		PID:   PID,
// 		Logs:  []*monitor.PageLog{},
// 	}

// }
