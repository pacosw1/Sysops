package system

import (
	"fmt"
	"sysops/globals"
	"sysops/monitor"
	"sysops/reader"
	"sysops/replacement"
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
	CommandNum   int
	Commands     []*monitor.CommandEvent
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
		CommandNum:   0,
		TimeStep:     0.0,
	}
}

//Start starts processing inputs from queue when avaliable
func (m *MemoryManager) Start() {
	m.Running = true
	m.processInputs()

}

//handleInput handles input received from reader and executes them
func (m *MemoryManager) handleInput(r monitor.Request) {

	// println("processing input")
	switch r.Type {

	case globals.Access:
		//initialize command logger
		m.Monitor.AddRequest(monitor.NewCommandEvent(globals.Access, m.CommandNum, m.TimeStep))
		PID := r
		vAddr := r.Args[1]
		mod := r.Args[2]
		m.AccessMemory(PID, vAddr, mod)
		break
	case globals.LoadP: //handles loading a new process

		m.Monitor.AddRequest(monitor.NewCommandEvent(globals.LoadP, m.CommandNum, m.TimeStep))
		size := r.Args[0]
		pID := r.Args[1]
		m.LoadProcess(types.NewProcess(pID, size, m.PageSize))
		break
	case globals.Print: //handles printing a message
		PrintMessage(r.Message, globals.Print)
		break
	case globals.FreeP: //handles freeing process vars

		//initalize command logger
		m.Monitor.AddRequest(monitor.NewCommandEvent(globals.FreeP, m.CommandNum, m.TimeStep))
		PID := r.Args[0]
		m.FreeProcess(PID)
		// m.Physical.View()
		break
	case globals.Stats:
		m.Pause()
		//display stats
		m.Resume()
		break
	case globals.End:
		m.Pause()
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
