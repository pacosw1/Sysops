package system

import (
	"fmt"
	"sysops/globals"
	"sysops/monitor"
	"sysops/types"
)

//InsertPage inserts a single page to Physicalory
func (m *MemoryManager) InsertPage(page *types.Page) {

	logger := m.Monitor.Requests[m.CommandNum]

	// replaceQ := m.ReplacementQ
	spaceTracker := m.Physical.SpaceTracker
	realMem := m.Physical.Memory

	//swap one frame to make space for new one
	if spaceTracker.Empty() {
		m.SwapOut()
	}

	before := types.CopyPage(page)

	//if space in physical memory
	if !spaceTracker.Empty() {

		//insertion takes 1 second per page

		availableFrame := spaceTracker.Pop()
		//update page
		page.PageFrame = availableFrame
		page.SwapFrame = -1

		//insert it to memory
		realMem[availableFrame] = page
		//add new page to replacement q
		m.ReplacementQ.Push(page)

	} else { //swap was full too
		fmt.Println("no space in memory")
	}

	after := types.CopyPage(page)
	log := monitor.NewPageLog(monitor.Insert, monitor.FromNew, monitor.ToMem, before, after, m.TimeStep)
	logger.AddLog(log)
	m.Monitor.AddLog(log)

	m.TimeStep++

}

//SwapIn swaps a single page to real memory
func (m *MemoryManager) SwapIn(page *types.Page) {

	logger := m.Monitor.Requests[m.CommandNum]

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

	log := monitor.NewPageLog(monitor.SwapIn, monitor.FromSwap, monitor.ToMem, before, after, m.TimeStep)
	logger.AddLog(log)
	m.Monitor.AddLog(log)

	m.TimeStep++

}

//SwapOut swaps a single page to swap memory
func (m *MemoryManager) SwapOut() {

	logger := m.Monitor.Requests[m.CommandNum]

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

	log := monitor.NewPageLog(monitor.SwapOut, monitor.FromMem, monitor.ToSwap, before, after, m.TimeStep)
	logger.AddLog(log)
	m.Monitor.AddLog(log)

	m.TimeStep++

}

//LoadProcess loads a process into memory
func (m *MemoryManager) LoadProcess(p *types.Process) {

	spaceTracker := m.Physical.SpaceTracker //keep track of free spaces
	swapSpaceTracker := m.Swap.SpaceTracker //keep track of swap frames

	//add to our list of processes
	m.ProcessList[p.PID] = p
	pages := len(p.Pages)

	//if not enough space, swap based on selected algo
	if pages > spaceTracker.Size() {

		dif := pages - spaceTracker.Size() // amount of free pages needed

		for i := 0; i < dif; i++ { //swap old pages to make space for new pages

			if !swapSpaceTracker.Empty() { //if we have free space inside Swap, swap oldest page
				m.SwapOut()
			} else {
				//to do, if Swap full do something
			}
		}
	}
	//add all pages to physical memory
	for i := 0; i < pages; i++ {
		page := p.Pages[i]
		m.InsertPage(page)
	}

	m.Monitor.Requests[m.CommandNum].End = m.TimeStep
	m.CommandNum++
}

func validSize(size int) bool {

	return size <= globals.MaxSize

}

//FreePages loads a proces unto memory

//PrintMessage prints message sent by input
func PrintMessage(msg []string, action string) {
	res := ""
	for _, word := range msg {
		res += word + " "
	}

	fmt.Println("INPUT: " + action)

	if len(res) == 0 {
		println("OUTPUT: Empty message \n\n")
	} else {
		fmt.Print("OUTPUT: " + res + "\n\n")
		fmt.Println("")
	}
}
