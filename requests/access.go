package requests

import (
	"fmt"
	"sysops/globals"
)

//AccessReq memory request and data storage
type AccessReq struct {
	Type  string
	Input []string
	PID   int
	VAddr int
	Mod   int
	Logs  []*PageLog
}

//NewAccessReq constructor
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

//GenerateOutput generates output for this request
func (f *AccessReq) GenerateOutput(output string) {

	fmt.Println("\n")

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	///ejecuta

	fmt.Println(output)

	logs := f.Logs

	if len(logs) > 0 {

		fmt.Println("\nFrames Swapped: \n")

		for i := 0; i < len(logs); i++ {

			log := logs[i]
			logType := log.Type

			if logType == SwapOut {
				fmt.Println("PID: %d  VP: %d Swap: %d \t", log.PageBefore.PID, log.PageBefore.ID, log.PageAfter.SwapFrame)
			}

		}
	}

	fmt.Println("\n")

}

//Args returns input arguments
func (f *AccessReq) Args() []int {
	args := []int{f.PID, f.VAddr, f.Mod}
	return args
}

//Type gets the request type
func (f *AccessReq) GetType() string {
	return f.Type
}

//GetLogs Return logs for this command
func (f *AccessReq) GetLogs() []*PageLog {
	return f.Logs
}

//AddLog adds a log to this command structure
func (f *AccessReq) AddLog(pageLog *PageLog) {
	f.Logs = append(f.Logs, pageLog)
}
