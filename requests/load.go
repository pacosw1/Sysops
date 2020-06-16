package requests

import (
	"fmt"
	"sysops/globals"
)

//FreeMem request to free mem
type LoadP struct {
	Type  string
	Input []string
	Size  int
	PID   int
	Logs  []*PageLog
}

//NewFreeMemReq n
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

//GenerateOutput p
func (f *LoadP) GenerateOutput() {
	fmt.Println("\n")

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	fmt.Println(f.Input)

	///ejecuta
	logs := f.Logs

	//display assigned frames in memory
	fmt.Println("Assigned: \n")
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

	fmt.Println("-------------------------------------------------------------------\n\n")

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

//GetLogs t
func (f *LoadP) GetLogs() []*PageLog {
	return f.Logs
}
