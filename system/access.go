package system

import (
	"fmt"
)

func getRealAddr(realPage, pageSize, offset int) int {
	return (realPage * pageSize) + offset
}

//AccessMemory access or modify bits inside pages
func (mm *MemoryManager) AccessMemory(PID int, vAddr int, m int) {

	p := mm.ProcessList[PID]

	if p == nil {
		fmt.Printf("Process does not exist in simulation")
		return
	}

	//validation bit overflow
	if vAddr >= p.Size {
		return
	}

	//get offset and virtual page for bit
	info := p.GetInfo(vAddr)
	page := p.Pages[info.Page]

	//if bit in real memory
	if page.SwapFrame >= 0 { //esta en el swap
		mm.SwapIn(page) //add it to physical memory
	}

	//if page modified
	if m == 1 {
		page.Mod = true
	}

	//update LRU if active
	mm.ReplacementQ.Push(page)

	// physicalAddress := getRealAddr(page.PageFrame, mm.PageSize, info.Offset)

	//records end of command
	mm.Monitor.Requests[mm.CommandNum].End = mm.TimeStep
	mm.CommandNum++

	mm.TimeStep += 0.1 // access time

	// fmt.Printf(" \nphysical Address: %d  PID: %d  ID: %d \n", physicalAddress, page.PID, page.ID)

}
