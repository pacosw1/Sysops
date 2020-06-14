package system

import (
	"fmt"
	"sysops/types"
)

func (m *MemoryManager) InsertPage(page *types.Page) {

}

func (m *MemoryManager) SwapPage(page *types.Page) {

	//if page already in swap, return
	if !page.InMem {
		return
	}

	//if page in physical memory swap it

}

//LoadProcess loads a proces into memory
func (m *MemoryManager) LoadProcess(p *types.Process) {

	realMem := m.PhysicalMem.Memory

	spaceTracker := m.PhysicalMem.SpaceTracker
	//add to our list of processes
	m.ProcessList[p.PID] = p
	pages := len(p.Pages)

	//if not enough space, swap based on selected algo
	if pages > spaceTracker.Size() {
		//FIFO

		fmt.Println("\n Time to Swap \n ")
		//amount of pages to be freed
		dif := pages - spaceTracker.Size()
		fmt.Printf("Replacing %d pages \n", dif)

		//swap old pages to make space for new pages
		for i := 0; i < dif; i++ {

			swapSpaceTracker := m.SwapMemory.SpaceTracker

			//get next page to be swapped in replacement algo
			oldestPage := m.ReplacementQ.Remove()
			fmt.Printf("Removing oldest Page with num %d and PID: %d \n\n", oldestPage.ID, oldestPage.PID)

			//if we have free space inside Swap
			if !swapSpaceTracker.Empty() {

				//get the next free slot in the list
				avaliableFrame := swapSpaceTracker.Pop()

				//free up space in physical memory and update FreePages list
				m.PhysicalMem.Memory[oldestPage.PageFrame] = nil
				swapSpaceTracker.Add(oldestPage.PageFrame)

				//update page details
				oldestPage.InMem = false
				oldestPage.PageFrame = -1
				oldestPage.SwapAddr = avaliableFrame

				//save page inside Swap
				m.SwapMemory.Memory[avaliableFrame] = oldestPage

				//update Swap Free Pages

			} else {
				//to do, if swapMemory full do something
			}

		}

		//add remaining pages to physical memory

		for i := 0; i < pages; i++ {

			//loading a page initially takes 1 second
			m.TimeStep++

			//get a free frame from physicalMemory
			avaliableFrame := spaceTracker.Pop()

			//update page in table
			page := p.Pages[i]

			page.InMem = true
			page.SwapAddr = -1
			page.InsertedAt = m.TimeStep
			page.PageFrame = avaliableFrame

			//insert page to memory, for locating purposes
			realMem[avaliableFrame] = page

			//add page to our replacement algorithm
			m.ReplacementQ.Add(page)

		}

		//LRU
	} else { //there is enough space to add pages

		//add all pages to physical memory
		for i := 0; i < pages; i++ {

			//loading a page initially takes 1 second
			m.TimeStep++

			//type assertion
			avaliableFrame := spaceTracker.Pop()

			//update page in table
			page := p.Pages[i]

			page.InMem = true
			page.InsertedAt = m.TimeStep
			page.PageFrame = avaliableFrame

			//insert page to memory, for locating purposes
			realMem[avaliableFrame] = page

			//add page to our replacement algorithm
			m.ReplacementQ.Add(page)
		}
	}
	//look for spaces in physical memory
	//fill in process pages in free slots and update Page table

	//if no spaces avaliable, swap items based on LRU or FIFO
	//do this until new program is inside memory

}
