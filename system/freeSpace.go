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

	//loop over page table
	for _, page := range pageTable {

		//if page in physical memory
		fmt.Printf("PID %d page# %d is in main memory, deleting...\n", page.PID, page.ID)
		if page.PageFrame != -1 {
			//free up memory
			temp := m.PhysicalMem.Memory[page.PageFrame]
			m.PhysicalMem.Memory[page.PageFrame] = nil

			//delete index from FreePages list
			dummy := m.SwapMemory.FreePages.Front()
			for dummy != nil {
				if dummy.Value == temp.PageFrame {
					m.SwapMemory.FreePages.Remove(dummy)
					break
				}
				dummy = dummy.Next()
			}

		} else { //else it's in swap memory
			fmt.Printf("PID %d page# %d is in swap memory, deleting...\n", page.PID, page.ID)

			temp := m.SwapMemory.Memory[page.SwapAddr]
			m.SwapMemory.Memory[page.SwapAddr] = nil

			dummy := m.SwapMemory.FreePages.Front()

			//delete index from FreePages list
			for dummy != nil {
				if dummy.Value == temp.SwapAddr {
					m.SwapMemory.FreePages.Remove(dummy)
					break
				}
				dummy = dummy.Next()
			}
			fmt.Printf("PID %d page# %d successfully deleted\n", page.PID, page.ID)
		}
	}

}
