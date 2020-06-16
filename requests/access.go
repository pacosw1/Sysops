package requests

import (
	"fmt"
	"sysops/globals"
)

//AccessReq request to free mem
type AccessReq struct {
	Type  string
	Input []string
	PID   int
	VAddr int
	Mod   int
	Logs  []*PageLog
}

//NewAccessReq n
func NewAccessReq(input []string, PID, vAddr, mod int) *AccessReq {

	return &AccessReq{
		Type:  globals.Access,
		Input: input,
		PID:   PID,
		VAddr: vAddr,
		Mod:   mod,
		Logs:  []*PageLog{},
	}
}

//GenerateOutput p
func (f *AccessReq) GenerateOutput(output string) {

	fmt.Println("\n")

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	///ejecuta

	logs := f.Logs

	fmt.Println("\nFrames Swapped: \n")
	for i := 0; i < len(logs); i++ {

		log := logs[i]
		logType := log.Type

		if logType == SwapOut {
			fmt.Println("PID: %d  VP: %d SWAPage: %d \t", log.PageBefore.PID, log.PageBefore.ID, log.PageAfter.SwapFrame)
		}

	}

	fmt.Println("\n")

	fmt.Println("-------------------------------------------------------------------\n\n")

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

//GetLogs t
func (f *AccessReq) GetLogs() []*PageLog {
	return f.Logs
}

//AddLog adds a log to the queue
func (f *AccessReq) AddLog(pageLog *PageLog) {
	f.Logs = append(f.Logs, pageLog)
}
