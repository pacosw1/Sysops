package simulation

import (
	"sysops/requests"
	"sysops/types"
)

//SwapIn swaps a single page to real memory
func (m *MemoryManager) SwapIn(page *types.Page) {

	monitor := m.Monitor
	req := monitor.Requests[m.Monitor.ReqNum]

	//keep track of free spaces in swap and real mem
	spaceTracker := m.Physical.SpaceTracker
	swapSpaceTracker := m.Swap.SpaceTracker

	//reference memories
	realMem := m.Physical.Memory
	swapMem := m.Swap.Memory

	//swapping takes one second per page

	//add it to free page list

	if spaceTracker.Empty() { //swap out a page is memory is full
		m.SwapOut()
	}

	swapSpaceTracker.Add(page.SwapFrame) //add released frame to free frames

	before := types.CopyPage(page)

	availableFrame := spaceTracker.Pop() //find a free space in real memory

	swapMem[page.SwapFrame] = nil //free page from swap memory

	//update page info
	page.SwapFrame = -1
	page.PageFrame = availableFrame

	//store it in real memory
	realMem[availableFrame] = page

	after := types.CopyPage(page)

	log := requests.NewPageLog(requests.SwapIn, requests.FromSwap, requests.ToMem, before, after, m.TimeStep)
	monitor.AddLog(log)
	req.AddLog(log)

	m.TimeStep++

}

//SwapOut swaps a single page to swap memory
func (m *MemoryManager) SwapOut() {

	monitor := m.Monitor
	req := monitor.Requests[m.Monitor.ReqNum]

	//if nothing to replace return
	if m.ReplacementQ.Empty() {
		return
	}
	//swapping takes 1 second per page

	page := m.ReplacementQ.Pop() //get oldest value to be replaced
	before := types.CopyPage(page)

	//keep track of free spaces in swap and real mem
	spaceTracker := m.Physical.SpaceTracker
	swapSpaceTracker := m.Swap.SpaceTracker

	//reference memories
	realMem := m.Physical.Memory
	swapMem := m.Swap.Memory

	//add it to free page list
	spaceTracker.Add(page.PageFrame) //add released frame to free frames

	realMem[page.PageFrame] = nil            //free page from real memory
	availableFrame := swapSpaceTracker.Pop() //update free frames

	//update page info
	page.PageFrame = -1
	page.SwapFrame = availableFrame

	//store it in swap
	swapMem[availableFrame] = page

	after := types.CopyPage(page)

	log := requests.NewPageLog(requests.SwapOut, requests.FromMem, requests.ToSwap, before, after, m.TimeStep)
	//save logs
	monitor.AddLog(log)
	req.AddLog(log)

	m.TimeStep++

}
