package system

import (
	"fmt"
	"sysops/globals"
	"sysops/types"
)

//InsertPage inserts a single page to Physicalory
func (m *MemoryManager) InsertPage(page *types.Page) *types.Page {

	// replaceQ := m.ReplacementQ
	spaceTracker := m.Physical.SpaceTracker
	realMem := m.Physical.Memory

	//swap one frame to make space for new one
	if spaceTracker.Empty() {
		m.SwapPage()
	}

	//if space in physical memory
	if !spaceTracker.Empty() {

		availableFrame := spaceTracker.Pop()
		// fmt.Printf("Inserting page at frame %d \n", availableFrame)
		//update page
		page.PageFrame = availableFrame
		page.SwapFrame = -1

		// fmt.Println(availableFrame)
		//insert it to memory
		realMem[availableFrame] = page
		//add new page to replacement q
		m.ReplacementQ.Add(page)

	} else { //swap was full too
		fmt.Println("no space in memory")
	}
	return page

}

//SwapPage swaps a single page to swap memory
func (m *MemoryManager) SwapPage() {

	//if nothing to replace return
	if m.ReplacementQ.Empty() {
		fmt.Println("No page to replace, empty memory")
		return
	}

	page := m.ReplacementQ.Remove() //get oldest value to be replaced

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

}

//LoadProcess loads a proces into memory
func (m *MemoryManager) LoadProcess(p *types.Process) {

	spaceTracker := m.Physical.SpaceTracker //keep track of free spaces

	//add to our list of processes
	m.ProcessList[p.PID] = p
	pages := len(p.Pages)

	//if not enough space, swap based on selected algo
	if pages > spaceTracker.Size() {

		dif := pages - spaceTracker.Size() // amount of free pages needed

		for i := 0; i < dif; i++ { //swap old pages to make space for new pages

			swapSpaceTracker := m.Swap.SpaceTracker //keep track of swap frames

			if !swapSpaceTracker.Empty() { //if we have free space inside Swap, swap oldest page
				m.SwapPage()
			} else {
				//to do, if Swap full do something
			}

		}

		for i := 0; i < pages; i++ { //add remaining pages to physical memory

			//loading a page initially takes 1 second
			m.TimeStep++
			page := p.Pages[i]
			m.InsertPage(page)

		}
	} else { //there is enough space to add pages

		//add all pages to physical memory
		for i := 0; i < pages; i++ {
			//loading a page initially takes 1 second
			m.TimeStep++
			page := p.Pages[i]
			m.InsertPage(page)

		}
	}
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
