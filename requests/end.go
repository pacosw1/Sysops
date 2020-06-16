package requests

import (
	"sysops/globals"
)

//EndReq request to free mem
type EndReq struct {
	Type  string
	Input []string
}

//NewEndReq n
func NewEndReq(input []string) *EndReq {

	return &EndReq{
		Type:  globals.End,
		Input: input,
	}
}

//GenerateOutput p
func (f *EndReq) GenerateOutput() {

}

//AddLog adds a log to the queue
func (f *EndReq) AddLog(pageLog *PageLog) {

}

//Args a
func (f *EndReq) Args() []int {
	return []int{}
}

//GetType t
func (f *EndReq) GetType() string {
	return f.Type
}
