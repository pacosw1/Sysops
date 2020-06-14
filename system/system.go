package system

import (
	"fmt"
	"sysops/globals"
	"sysops/input"
	"sysops/types"
	"sysops/virtual"
)

//MemoryManager manages and stores memory
type MemoryManager struct {
	PageSize     int
	Running      bool
	Reader       *input.Reader
	PhysicalMem  *virtual.Storage //physical mem
	SwapMemory   *virtual.Storage //swap mem
	ProcessList  []*types.Process //list of process objects
	CommandQueue chan string
	TimeStep     int
}

//New creates a new Memory Manager
func New(physicalSize, swapSize, pagesize int) *MemoryManager {
	return &MemoryManager{
		Running:     false,
		PageSize:    pagesize,
		Reader:      input.NewReader(),
		PhysicalMem: virtual.NewStorage(physicalSize, pagesize),
		SwapMemory:  virtual.NewStorage(swapSize, pagesize),
		ProcessList: make([]*types.Process, 0),
		TimeStep:    0,
	}
}

//Start starts processing inputs from queue when avaliable
func (m *MemoryManager) Start() {
	m.Running = true
	m.processInputs()

}

//handleInput handles input received from reader and executes them
func (m *MemoryManager) handleInput(r *types.Request) {

	println("processing input")
	switch r.Type {
	case globals.Access:
		break
	case globals.LoadP: //handles loading a new process
		size := r.Args[0]
		pID := r.Args[1]
		m.LoadProcess(types.NewProcess(pID, size, m.PageSize))
		break
	case globals.Print: //handles printing a message
		PrintMessage(r.Message, globals.Print)
		break
	case globals.FreeP: //handles freeing process vars
		break
	case globals.Stats:
		m.Pause()
		//display stats
		m.Resume()
		break
	case globals.End:
		m.Pause()
		m.PhysicalMem.View()
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
