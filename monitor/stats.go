package monitor

import (
	"fmt"
	"math"
	"sysops/requests"
)

//Monitor Stats
type Monitor struct {
	Logs     []*requests.PageLog
	Requests []requests.Request
	ProStats map[int]*ProStats
	ReqNum   int
}

//NewMonitor new monitor
func NewMonitor() *Monitor {
	return &Monitor{
		Logs:     []*requests.PageLog{},
		Requests: []requests.Request{},
		ProStats: make(map[int]*ProStats, 0),
		ReqNum:   0,
	}
}

func (m *Monitor) Reset() {
	m.Logs = []*requests.PageLog{}
	m.Requests = []requests.Request{}
	m.ProStats = map[int]*ProStats{}
	m.ReqNum = 0
}

//AddLog adds a new request to be logged
func (m *Monitor) AddLog(c *requests.PageLog) {
	m.Logs = append(m.Logs, c)
}

//AddRequest adds a new request to be logged
func (m *Monitor) AddRequest(r requests.Request) {
	m.Requests = append(m.Requests, r)
}

//GenerateStats generates stats up to this point
func (m *Monitor) GenerateStats(timestamp float32) {

	logs := m.Logs
	proStats := m.ProStats

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println("STATISTICS FOR THIS ROUND\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	//total swaps with input
	swapIns := 0
	swapOuts := 0

	for i := 0; i < len(logs); i++ {

		log := logs[i]

		switch log.Type {
		case requests.SwapIn: //general
			swapIns++
			break
		case requests.SwapOut: //general
			swapOuts++
			break
		case requests.Freed: //process specific
			p := proStats[log.PageBefore.PID]
			// PID := log.PageBefore.PID
			// fmt.Printf("TimeStamp: %f", p.EndStep)
			p.EndStep = float32(math.Max(float64(p.EndStep), float64(log.TimeStep)))
			break
		case requests.PageFault:
			proStats[log.PageAfter.PID].PageFaults++
			break
		default:
			break
		}
	}

	//print process stats
	turnAroundSum := 0.0

	for PID, stats := range proStats {
		turnAround := stats.EndStep
		if stats.EndStep >= stats.StartStep {
			turnAroundSum += float64(turnAround)
		}
		fmt.Printf("ProcessID (PID): %d \nPageFaults: %d \n", PID, stats.PageFaults)
		if stats.EndStep == -1 {
			fmt.Printf("Turnaround:\t No termino\n\n")
		} else {
			fmt.Printf("Turnaround: %f \t \n\n", turnAround)

		}
	}

	fmt.Printf("\nSwap IN: %d \t Swap OUT: %d \n\n", swapIns, swapOuts)
	prom := turnAroundSum / float64(len(proStats))
	fmt.Printf("Turnaround Promedio: %f Timestamp: %f \n\n", (prom), timestamp)

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println("END OF STATISTCS\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

}
