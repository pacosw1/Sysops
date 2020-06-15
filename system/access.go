package system

import (
	"fmt"
	"sysops/types"
)

func getRealAddr(realPage, pageSize, offset int) int {
	return (realPage * pageSize) + offset
}

//AccessMemory access or modify bits inside pages
func (mm *MemoryManager) AccessMemory(p *types.Process, vAddr int, m int) {

	//validation bit overflow
	if vAddr >= p.Size {
		return
	}

	//get offset and virtual page for bit
	info := p.GetInfo(vAddr)
	page := p.Pages[info.Page]

	//if bit in real memory
	if page.SwapFrame >= 0 { //esta en el swap
		fmt.Printf("position in swap %d \n ", page.SwapFrame)
		mm.FreePage(page)
		mm.InsertPage(page)

	}

	if m == 1 {
		page.Mod = true
		//Actualizar LRU
	}

	physicalAddress := getRealAddr(page.PageFrame, mm.PageSize, info.Offset)
	fmt.Printf("realAddr: %d \n ", physicalAddress)

}
