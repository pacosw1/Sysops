package system

import (
	"fmt"
	"sysops/types"
)

//FreePages frees all pages
func (m *MemoryManager) FreePages(p *types.Process) {

	process := m.ProcessList[p.PID]

	//if process does not exist
	if process == nil {
		fmt.Printf("Process %d does not exist", p.PID)
		return
	}

	//if it does exist
	pageTable := process.Pages

	//keep track of available frames in each memory
	physicalSpaceTracker := m.PhysicalMem.SpaceTracker
	swapSpaceTracker := m.SwapMemory.SpaceTracker

	//loop over page table
	for _, page := range pageTable {

		//if page in physical memory
		fmt.Printf("PID %d page# %d is in main memory, deleting...\n", page.PID, page.ID)
		if page.PageFrame != -1 {
			//free up memory
			m.PhysicalMem.Memory[page.PageFrame] = nil
			physicalSpaceTracker.Remove(page.PageFrame)

			//delete index from FreePages list

		} else { //else it's in swap memory
			fmt.Printf("PID %d page# %d is in swap memory, deleting...\n", page.PID, page.ID)

			m.SwapMemory.Memory[page.SwapAddr] = nil
			swapSpaceTracker.Remove(page.SwapAddr)
			//delete index from FreePages list

			fmt.Printf("PID %d page# %d successfully deleted\n", page.PID, page.ID)
		}
	}

}
