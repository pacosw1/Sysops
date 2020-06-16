package requests
package requests

import (
	"sysops/globals"
)

//StatsReq request to free mem
type StatsReq struct {
	Type  string
	Input []string
}

//NewStatsReq n
func NewStatsReq(input []string) *StatsReq {

	return &StatsReq{
		Type:  globals.End,
		Input: input,
	}
}

//GenerateOutput p
func (f *StatsReq) GenerateOutput() {

	

}

//AddLog adds a log to the queue
func (f *StatsReq) AddLog(pageLog *PageLog) {

}

//Args a
func (f *StatsReq) Args() []int {
	return []int{}
}

//GetType t
func (f *StatsReq) GetType() string {
	return f.Type
}
