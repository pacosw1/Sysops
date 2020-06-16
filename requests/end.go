package requests

import (
	"sysops/globals"
)

//EndReq request to end program
type EndReq struct {
	Type  string
	Input []string
}

//NewEndReq constructor
func NewEndReq(input []string) *EndReq {

	return &EndReq{
		Type:  globals.End,
		Input: input,
	}
}

//GenerateOutput Generates output for this command
func (f *EndReq) GenerateOutput(output string) {

}

//AddLog adds a log to the queue
func (f *EndReq) AddLog(pageLog *PageLog) {

}

//Args input arguments
func (f *EndReq) Args() []int {
	return []int{}
}

//GetType returns request type
func (f *EndReq) GetType() string {
	return f.Type
}
