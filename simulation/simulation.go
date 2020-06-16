package simulation

import (
	"fmt"
	"sysops/globals"
	"sysops/monitor"
	"sysops/reader"
	"sysops/replacement"
	"sysops/requests"
	"sysops/types"
	"sysops/virtual"
)

//MemoryManager manages and stores memory
type MemoryManager struct {
	PageSize     int
	Running      bool
	Reader       *reader.Reader
	Physical     *virtual.Storage       //physical mem
	Swap         *virtual.Storage       //swap mem
	ProcessList  map[int]*types.Process //list of process objects
	ReplacementQ replacement.Algo       //replacement algorithm, FIFO or LRU
	Monitor      *monitor.Monitor
	TimeStep     float32
}

//New creates a new Memory Manager
func New(physicalSize, swapSize, replacementAlgo int, pagesize int) *MemoryManager {

	var replacementQ replacement.Algo

	//select replacement algo interface, switch open to future additions
	switch replacementAlgo {
	case globals.FIFO:
		replacementQ = replacement.NewFIFO()
		break
	case globals.LRU:
		replacementQ = replacement.NewLRU()
		break
	}

	return &MemoryManager{
		Running:      false,
		PageSize:     pagesize,
		Reader:       reader.NewReader(),
		Physical:     virtual.NewStorage(physicalSize, pagesize),
		Swap:         virtual.NewStorage(swapSize, pagesize),
		ProcessList:  map[int]*types.Process{},
		Monitor:      monitor.NewMonitor(),
		ReplacementQ: replacementQ,
		TimeStep:     0.0,
	}
}

//Start starts processing inputs from queue when avaliable
func (m *MemoryManager) Start() {
	m.Running = true
	m.processInputs()

}

//handleInput handles input received from reader and executes them
func (m *MemoryManager) handleInput(r requests.Request) {
	monitor := m.Monitor
	// fmt.Println(r.Args())
	tipe := r.GetType()
	switch tipe {
	case globals.Access:
		//initialize command logger
		m.Monitor.AddRequest(r)
		args := r.Args()

		PID := args[0]
		vAddr := args[1]
		mod := args[2]

		m.AccessMemory(PID, vAddr, mod)
		monitor.Requests[monitor.ReqNum].GenerateOutput()
		monitor.ReqNum++
		break
	case globals.LoadP: //handles loading a new process

		m.Monitor.AddRequest(r)

		args := r.Args()
		PID := args[0]
		size := args[1]

		m.LoadProcess(types.NewProcess(PID, size, m.PageSize))
		monitor.Requests[monitor.ReqNum].GenerateOutput()
		monitor.ReqNum++

		break
	case globals.Print: //handles printing a message

		monitor.AddRequest(r)
		monitor.Requests[monitor.ReqNum].GenerateOutput()
		monitor.ReqNum++

		break
	case globals.FreeP: //handles freeing process vars

		m.Monitor.AddRequest(r)
		args := r.Args()

		PID := args[0]
		m.FreeProcess(PID)

		monitor.Requests[monitor.ReqNum].GenerateOutput()
		monitor.ReqNum++
		break
	case globals.Stats:
		m.Pause()
		//display stats
		m.Resume()
		monitor.ReqNum++

		break
	case globals.End:
		monitor.AddRequest(r)
		monitor.ReqNum++
		m.Pause()
		// m.Physical.View()
		// m.Swap.View()
		break
	default:
		fmt.Println("Command could not be read ")
	}
}

func (m *MemoryManager) processInputs() {
	for m.Running {
		//receive a request from the queue
		req := <-m.Reader.CommandQ
		m.handleInput(req)
	}
}

//Pause pauses input processing
func (m *MemoryManager) Pause() {
	m.Running = false
}

//Resume resumes processing inputs
func (m *MemoryManager) Resume() {
	m.Running = true
	m.processInputs()
}
