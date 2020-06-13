package system

import (
	"sysops/types"
	"sysops/virtual"
)

//MemoryManager manages and stores memory
type MemoryManager struct {
	PageSize     int
	Running      bool
	PhysicalMem  *virtual.Storage
	SwapMemory   *virtual.Storage
	ProcessList  []*types.Process
	CommandQueue chan string
	TimeStep     int
}

//New creates a new Memory Manager
func New(physicalSize, swapSize, pagesize int) *MemoryManager {
	return &MemoryManager{
		Running:      false,
		PageSize:     pagesize,
		PhysicalMem:  virtual.NewStorage(physicalSize, pagesize),
		SwapMemory:   virtual.NewStorage(swapSize, pagesize),
		CommandQueue: make(chan string, 1000),
		ProcessList:  make([]*types.Process, 0),
		TimeStep:     0,
	}
}
