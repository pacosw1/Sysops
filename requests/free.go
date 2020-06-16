package requests

import (
	"fmt"
	"sysops/globals"
)

//FreeMem request to free mem
type FreeMem struct {
	Type  string
	Input []string
	PID   int
	Logs  []*PageLog
}

//NewFreeMemReq constructor
func NewFreeMemReq(input []string, PID int) *FreeMem {

	return &FreeMem{
		Type:  globals.FreeP,
		Input: input,
		PID:   PID,
		Logs:  []*PageLog{},
	}
}

//GenerateOutput generates output based on logs and function
func (f *FreeMem) GenerateOutput(output string) {

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "Start\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	///ejecuta

	fmt.Println(output)

	logs := f.Logs

	if len(logs) == 0 {

	} else {

		fmt.Println("\nFrames Freed: \n")
		for i := 0; i < len(logs); i++ {

			log := logs[i]
			logType := log.Type
			if logType == Freed {

				if log.PageBefore.PageFrame != -1 {
					fmt.Printf("PID: %d  VP: %d  Frame(Memory): %d \t\t", log.PageAfter.PID, log.PageBefore.ID, log.PageBefore.PageFrame)

				} else {
					fmt.Printf("PID: %d  VP: %d  Frame(Swap): %d \t\t", log.PageBefore.PID, log.PageBefore.ID, log.PageBefore.SwapFrame)

				}

			}

		}

		fmt.Println("\n")

	}

}

//Args input arguments
func (f *FreeMem) Args() []int {
	args := []int{f.PID}
	return args
}

//Type get request type
func (f *FreeMem) GetType() string {
	return f.Type
}

//GetLogs get logs for this request
func (f *FreeMem) GetLogs() []*PageLog {
	return f.Logs
}

//AddLog adds a log to the queue
func (f *FreeMem) AddLog(pageLog *PageLog) {
	f.Logs = append(f.Logs, pageLog)
}
