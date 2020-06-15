package monitor

import (
	"fmt"
	"sysops/types"
)

//Stats Monitor
type Monitor struct {
	Logs     []*PageLog
	Requests []*CommandEvent
}

//NewMonitor new monitor
func NewMonitor() *Monitor {
	return &Monitor{
		Logs:     []*PageLog{},
		Requests: []*CommandEvent{},
	}
}

//AddRequest adds a new request to be logged
func (m *Monitor) AddRequest(c *CommandEvent) {
	m.Requests = append(m.Requests, c)
}

//Swap

const (
	//Insert a page was loaded into memory
	Insert = "INSERT"
	//SwapIn a page was loaded into memory from swap storage
	SwapIn = "SWAPIN"
	//SwapOut a page was swapped out of memory
	SwapOut = "SWAPOUT"
	//Freed a page was freed from memory
	Freed = "FREED"
	//PageFault a page fault is created to load a certain process into memory
	PageFault = "PAGEFUALT"
)

const (
	//FromNew coming from new process
	FromNew = "FROMNULL"
	//FromSwap coming from swap memory
	FromSwap = "FROMSWAP"
	//FromMem coming from real memory
	FromMem = "FROMMEM"

	//ToMem loaded to Memory
	ToMem = "TOMEM"
	//ToSwap loaded to swap storage
	ToSwap = "TOSWAP"
	//ToNull freed from memory (deleted)
	ToNull = "TONULL"
)

//PageLog logs page actions and data
type PageLog struct {
	Type        string
	Source      string
	Destination string
	PageBefore  *types.Page
	PageAfter   *types.Page
	TimeStep    float32
}

//AddLog adds a log to the queue
func (m *Monitor) AddLog(pageLog *PageLog) {
	m.Logs = append(m.Logs, pageLog)
}

//Print []
func (l *PageLog) Print() {

	t := l.Type
	src := l.Source
	dest := l.Destination

	fmt.Printf("Type: " + t + "\n")
	fmt.Printf("Src: " + src + "\n")
	fmt.Printf("Dest: " + dest + "\n\n")

}

//NewPageLog c
func NewPageLog(t string, src string, dest string, before, after *types.Page, step float32) *PageLog {

	return &PageLog{
		Type:        t,
		Source:      src,
		Destination: dest,
		PageBefore:  before,
		PageAfter:   after,
		TimeStep:    step,
	}

}

/*
switch t {
	case Insert:
		break
	case SwapIn:
		break
	case SwapOut:
		break
	case Freed:
		break
	case PageFault:
		break
	}
*/
//CommandEvent c
type CommandEvent struct {
	Type   string
	Start  float32
	End    float32
	Output []int
	Logs   []*PageLog
	ID     int
}

//NewCommandEvent c
func NewCommandEvent(t string, id int, start float32) *CommandEvent {
	return &CommandEvent{
		Type:   t,
		Start:  start,
		Logs:   []*PageLog{},
		Output: []int{},
		ID:     id,
	}
}

//AddLog appendsd a log to the current command that is running
func (c *CommandEvent) AddLog(p *PageLog) {

	c.Logs = append(c.Logs, p)
}

//GenerateReport r
func (m *Monitor) GenerateReport(id int) {

	//collect all logs in command timestamp

}
