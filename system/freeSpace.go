package system

import (
	"fmt"
	"sysops/types"
)

//FreeProcess frees all pages from a specific process
func (m *MemoryManager) FreeProcess(p *types.Process) {

	process := m.ProcessList[p.PID]

	//if process does not exist
	if process == nil {
		fmt.Printf("Process %d does not exist", p.PID)
		return
	}

	//if it does exist
	pageTable := process.Pages

	//loop over page table
	for _, page := range pageTable {
		m.FreePage(page)
	}

}

//FreePage frees a page from memory
func (m *MemoryManager) FreePage(p *types.Page) {

	if p.PageFrame >= 0 {
		pageFrame := p.PageFrame
		m.Physical.SpaceTracker.Add(pageFrame)
		m.Physical.Memory[pageFrame] = nil
		//update page just in case
		p.PageFrame = -1
		p.SwapFrame = -1
	} else {
		swapFrame := p.SwapFrame
		m.Swap.SpaceTracker.Add(swapFrame)
		m.Swap.Memory[swapFrame] = nil
		p.PageFrame = -1
		p.SwapFrame = -1
	}
}
