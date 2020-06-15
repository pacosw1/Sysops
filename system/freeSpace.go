package system

import (
	"fmt"
	"sysops/globals"
	"sysops/monitor"
	"sysops/types"
)

//FreeProcess frees all pages from a specific process
func (m *MemoryManager) FreeProcess(PID int) {

	//initalize command logger
	m.Monitor.AddRequest(monitor.NewCommandEvent(globals.FreeP, m.CommandNum, m.TimeStep))

	process := m.ProcessList[PID]

	//if process does not exist
	if process == nil {
		fmt.Printf("Process %d does not exist", PID)
		return
	}

	//if it does exist
	pageTable := process.Pages

	//loop over page table
	for _, page := range pageTable {
		m.FreePage(page)
	}

	//record command end
	m.Monitor.Requests[m.CommandNum].End = m.TimeStep
	m.CommandNum++

}

//FreePage frees a page from memory
func (m *MemoryManager) FreePage(p *types.Page) {

	logger := m.Monitor.Requests[m.CommandNum]

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

		log := monitor.NewPageLog(monitor.Freed, monitor.FromSwap, monitor.ToNull, before, after, m.TimeStep)
		logger.AddLog(log)
		m.Monitor.AddLog(log)

	} else {
		swapFrame := p.SwapFrame
		m.Swap.SpaceTracker.Add(swapFrame)
		m.Swap.Memory[swapFrame] = nil
		p.PageFrame = -1
		p.SwapFrame = -1
		after := types.CopyPage(p)

		log := monitor.NewPageLog(monitor.Freed, monitor.FromSwap, monitor.ToNull, before, after, m.TimeStep)
		logger.AddLog(log)
		m.Monitor.AddLog(log)

	}

	m.TimeStep += 0.1 //process time

}
