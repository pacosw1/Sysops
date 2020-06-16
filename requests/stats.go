package requests

import (
	"fmt"
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
		Type:  globals.Stats,
		Input: input,
	}
}

//GenerateOutput p
func (f *StatsReq) GenerateOutput(output string) {

	fmt.Println(output)

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
