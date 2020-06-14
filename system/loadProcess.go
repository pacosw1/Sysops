package system

import (
	"fmt"
	"sysops/types"
)

//LoadProcess loads a proces into memory
func (m *MemoryManager) LoadProcess(p *types.Process) {

	realMem := m.PhysicalMem.Memory
	freeSlots := m.PhysicalMem.FreePages.Front()

	//add to our list of processes
	m.ProcessList[p.PID] = p
	pages := len(p.Pages)

	//if not enough space, swap based on selected algo
	if pages > m.PhysicalMem.FreePages.Len() {
		//FIFO

		fmt.Println("\n Time to Swap \n ")
		//amount of pages to be freed
		dif := pages - m.PhysicalMem.FreePages.Len()
		fmt.Printf("Replacing %d pages \n", dif)

		//swap old pages to make space for new pages
		for i := 0; i < dif; i++ {

			//get next page to be swapped in replacement algo
			oldestPage := m.ReplacementQ.Remove()
			fmt.Printf("Removing oldest Page with num %d and PID: %d \n\n", oldestPage.ID, oldestPage.PID)

			//if we have free space inside Swap
			if m.SwapMemory.FreePages.Len() > 0 {

				//get the next free slot in the list
				freeSlot := m.SwapMemory.FreePages.Front()
				freeFrame, _ := freeSlot.Value.(int)

				//free up space in physical memory and update FreePages list
				m.PhysicalMem.Memory[oldestPage.PageFrame] = nil
				m.PhysicalMem.FreePages.PushBack(oldestPage.PageFrame)

				//update page details
				oldestPage.InMem = false
				oldestPage.PageFrame = -1
				oldestPage.SwapAddr = freeFrame

				//save page inside Swap
				m.SwapMemory.Memory[freeFrame] = oldestPage

				//update Swap Free Pages
				m.SwapMemory.FreePages.Remove(freeSlot)

			} else {
				//to do, if swapMemory full do something
			}

		}

		//add remaining pages to physical memory

		for i := 0; i < pages; i++ {

			//loading a page initially takes 1 second
			m.TimeStep++

			freeSlots := m.PhysicalMem.FreePages.Front()

			//type assertion
			pageNum, _ := freeSlots.Value.(int)

			//update page in table
			page := p.Pages[i]

			page.InMem = true
			page.SwapAddr = -1
			page.InsertedAt = m.TimeStep
			page.PageFrame = pageNum

			//insert page to memory, for locating purposes
			realMem[pageNum] = page

			//add page to our replacement algorithm
			m.ReplacementQ.Add(page)

			//next free slot
			next := freeSlots.Next()

			//remove value from list since it is no longer free
			m.PhysicalMem.FreePages.Remove(freeSlots)
			//update list
			freeSlots = next

		}

		//LRU
	} else { //there is enough space to add pages

		//add all pages to physical memory
		for i := 0; i < pages; i++ {

			//loading a page initially takes 1 second
			m.TimeStep++

			//type assertion
			pageNum, _ := freeSlots.Value.(int)

			//update page in table
			page := p.Pages[i]

			page.InMem = true
			page.InsertedAt = m.TimeStep
			page.PageFrame = pageNum

			//insert page to memory, for locating purposes
			realMem[pageNum] = page

			//add page to our replacement algorithm
			m.ReplacementQ.Add(page)

			//next free slot
			next := freeSlots.Next()

			//remove value from list since it is no longer free
			m.PhysicalMem.FreePages.Remove(freeSlots)
			//update list
			freeSlots = next
		}
	}
	//look for spaces in physical memory
	//fill in process pages in free slots and update Page table

	//if no spaces avaliable, swap items based on LRU or FIFO
	//do this until new program is inside memory

}
