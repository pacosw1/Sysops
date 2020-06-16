package simulation

import (
	"fmt"
	"sysops/requests"
	"sysops/types"
)

//FreeProcess frees all pages from a specific process
func (m *MemoryManager) FreeProcess(PID int) {

	process := m.ProcessList[PID]

	//if process does not exist
	if process == nil {
		fmt.Printf("Process %d does not exist \n", PID)
		return
	}

	//if it does exist
	pageTable := process.Pages

	//loop over page table
	for _, page := range pageTable {
		m.FreePage(page)
	}

	//record command end
	delete(m.ProcessList, PID)

}

//FreePage frees a page from memory
func (m *MemoryManager) FreePage(p *types.Page) {

	monitor := m.Monitor
	//current request
	req := monitor.Requests[m.Monitor.ReqNum]

	var before *types.Page = types.CopyPage(p)

	if p.PageFrame >= 0 {
		pageFrame := p.PageFrame
		m.Physical.SpaceTracker.Add(pageFrame)
		m.Physical.Memory[pageFrame] = nil
		m.ReplacementQ.Remove(p)
		//update page just in case
		p.PageFrame = -1
		p.SwapFrame = -1
		after := types.CopyPage(p)

		log := requests.NewPageLog(requests.Freed, requests.FromSwap, requests.ToNull, before, after, m.TimeStep)
		//save logs
		monitor.AddLog(log)
		req.AddLog(log)

	} else {
		swapFrame := p.SwapFrame
		m.Swap.SpaceTracker.Add(swapFrame)
		m.Swap.Memory[swapFrame] = nil
		p.PageFrame = -1
		p.SwapFrame = -1
		after := types.CopyPage(p)

		log := requests.NewPageLog(requests.Freed, requests.FromSwap, requests.ToNull, before, after, m.TimeStep)
		//save logs
		monitor.AddLog(log)
		req.AddLog(log)

	}

	m.TimeStep += 0.1 //process time

}
