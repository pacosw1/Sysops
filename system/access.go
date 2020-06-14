package system

// import (
// 	"fmt"
// 	"sysops/types"
// )

// //AccessMemory access or modify bits inside pages
// func (mm *MemoryManager) AccessMemory(p *types.Process, vAddr int, m int) {
// 	//bit info

// 	//6

// 	//validation
// 	if p.Size < vAddr {

// 		info := p.GetInfo(vAddr)
// 		page := p.Pages[info.Page]

// 		//esta en memoria real
// 		if page.PageFrame >= 0 {

// 		} else { //esta en el swap
// 			fmt.Printf("position in swap %d \n ", page.SwapAddr)

// 		}

// 		if m == 0 {
// 			physicalAddress := info.Page*mm.PageSize + info.Offset
// 			fmt.Printf("physical address %d \n ", physicalAddress)

// 		} else {

// 			page.Mod = true

// 		}
// 	}

// }
