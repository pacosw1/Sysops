package requests

import (
	"fmt"
	"sysops/globals"
)

//LoadP load process request
type LoadP struct {
	Type  string
	Input []string
	Size  int
	PID   int
	Logs  []*PageLog
}

//NewLoadP constructor
func NewLoadP(input []string, PID int, size int) *LoadP {

	return &LoadP{
		Type:  globals.LoadP,
		Input: input,
		PID:   PID,
		Size:  size,
		Logs:  []*PageLog{},
	}
}

//AddLog adds a log to the queue
func (f *LoadP) AddLog(pageLog *PageLog) {
	f.Logs = append(f.Logs, pageLog)
}

//GenerateOutput generates log based output and function
func (f *LoadP) GenerateOutput(output string) {

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	///ejecuta
	logs := f.Logs

	fmt.Println(output)

	if len(logs) == 0 {

	} else {
		//display assigned frames in memory
		fmt.Println("\nAssigned Frames: \n")
		for i := 0; i < len(logs); i++ {

			log := logs[i]
			logType := log.Type

			//display pages and their current location that were swapped out from real memory
			if logType == Insert {
				fmt.Printf("PID: %d  VP: %d  Frame: %d \t\t", log.PageAfter.PID, log.PageAfter.ID, log.PageAfter.PageFrame)

			}

		}

		//display swapped frames if any
		fmt.Println("\n\nFrames Swapped: \n")
		for i := 0; i < len(logs); i++ {

			log := logs[i]
			logType := log.Type

			//display pages and their current location that were swapped out from real memory
			if logType == SwapOut {
				fmt.Printf("PID: %d  VP: %d  Frame: %d \t\t", log.PageAfter.PID, log.PageAfter.ID, log.PageAfter.SwapFrame)

			}

		}

		fmt.Println("\n")

	}

}

//Args input arguments
func (f *LoadP) Args() []int {
	args := []int{f.PID, f.Size}
	return args
}

//GetType return request type
func (f *LoadP) GetType() string {
	return f.Type
}

//GetLogs getLogs
func (f *LoadP) GetLogs() []*PageLog {
	return f.Logs
}
