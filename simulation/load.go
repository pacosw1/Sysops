package simulation

import (
	"fmt"
	"sysops/globals"
	"sysops/monitor"
	"sysops/requests"
	"sysops/types"
)

//InsertPage inserts a single page to Physicalory
func (m *MemoryManager) InsertPage(page *types.Page) {

	monitor := m.Monitor
	req := monitor.Requests[m.Monitor.ReqNum]

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
	log := requests.NewPageLog(requests.Insert, requests.FromNew, requests.ToMem, before, after, m.TimeStep)

	//save logs
	monitor.AddLog(log)
	req.AddLog(log)

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
				//according to pdf, this causes a page fault
				m.Monitor.AddLog(requests.NewPageLog(requests.PageFault, requests.FromNew, requests.ToMem, &types.Page{PID: p.PID}, &types.Page{PID: p.PID}, m.TimeStep))
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

	m.Monitor.ProStats[p.PID] = monitor.NewProStats(m.TimeStep)

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
