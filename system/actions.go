package system

import (
	"fmt"
	"sysops/types"
)

//LoadProcess loads a proces unto memory
func (m *MemoryManager) LoadProcess(p *types.Process) {

	realMem := m.PhysicalMem.Memory
	freeSlots := m.PhysicalMem.FreePages.Front()

	//add to our list of processes
	m.ProcessList = append(m.ProcessList, p)
	pages := len(p.Pages)

	//if not enough space, swap based on selected algo
	if pages > m.PhysicalMem.FreePages.Len() {
		println("Time to swap")
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

//AccessMemory access or modify bits inside pages
func (m *MemoryManager) AccessMemory(p *types.Process) {

}

//FreePages loads a proces unto memory
func (m *MemoryManager) FreePages(p *types.Process) {

	//look for space inside swap memory

	//preallocate space in swap

	//look for spaces in physical memory
	//fill in process pages in free slots and update Page table

	//if no spaces avaliable, swap items based on LRU or FIFO
	//do this until new program is inside memory

}

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
