package simulation

import (
	"strconv"
	"sysops/requests"
	"sysops/types"
)

func getRealAddr(realPage, pageSize, offset int) int {
	return (realPage * pageSize) + offset
}

//AccessMemory access or modify bits inside pages
func (mm *MemoryManager) AccessMemory(PID int, vAddr int, m int) string {

	p := mm.ProcessList[PID]

	output := ""

	//Process doesn not exist
	if p == nil {
		return ("Process does not exist in simulation")
	}

	//validation bit overflow
	if vAddr >= p.Size {
		return "Invalid virtual address, out of range!"
	}

	//get offset and virtual page for bit
	info := p.GetInfo(vAddr)
	page := p.Pages[info.Page]

	//page initial state
	before := types.CopyPage(page)

	//if bit in real memory
	if page.SwapFrame >= 0 { //esta en el swap
		output += "Process in Swap Frame: " + strconv.Itoa(page.SwapFrame)

		mm.SwapIn(page) //add it to physical memory

		//page state after swap
		after := types.CopyPage(page)

		//create a page fault log //PAGE FAULTS ONLY OCCUR WHEN ACCESSING MEMORY STORED IN SWAP
		mm.Monitor.AddLog(requests.NewPageLog(requests.PageFault, requests.FromSwap, requests.ToMem, before, after, mm.TimeStep))

	}

	//if page modified
	if m == 1 {
		page.Mod = true
	}

	//update LRU if active
	mm.ReplacementQ.Push(page)

	physicalAddress := getRealAddr(page.PageFrame, mm.PageSize, info.Offset)

	output += "Physical Address for Virtual Adress " + strconv.Itoa(vAddr) + " is now located at " + strconv.Itoa(physicalAddress)

	mm.TimeStep += 0.1 // access time

	return output

}
