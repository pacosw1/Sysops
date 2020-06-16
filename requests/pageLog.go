package requests

import (
	"fmt"
	"sysops/types"
)

const (
	//Insert a page was loaded into memory
	Insert = "INSERT"
	//SwapIn a page was loaded into memory from swap storage
	SwapIn = "SWAPIN"
	//SwapOut a page was swapped out of memory
	SwapOut = "SWAPOUT"
	//Freed a page was freed from memory
	Freed = "FREED"
	//PageFault a page fault is created to load a certain process into memory
	PageFault = "PAGEFAULT"
)

const (
	//FromNew coming from new process
	FromNew = "FROMNULL"
	//FromSwap coming from swap memory
	FromSwap = "FROMSWAP"
	//FromMem coming from real memory
	FromMem = "FROMMEM"

	//ToMem loaded to Memory
	ToMem = "TOMEM"
	//ToSwap loaded to swap storage
	ToSwap = "TOSWAP"
	//ToNull freed from memory (deleted)
	ToNull = "TONULL"
)

//Page log used to generate stats
type PageLog struct {
	Type        string
	Source      string
	Destination string
	PageBefore  *types.Page
	PageAfter   *types.Page
	TimeStep    float32
}

//Print duh
func (l *PageLog) Print() {

	t := l.Type
	src := l.Source
	dest := l.Destination

	fmt.Printf("Type: " + t + "\n")
	fmt.Printf("Src: " + src + "\n")
	fmt.Printf("Dest: " + dest + "\n\n")

}

//NewPageLog constructor
func NewPageLog(t string, src string, dest string, before, after *types.Page, step float32) *PageLog {

	return &PageLog{
		Type:        t,
		Source:      src,
		Destination: dest,
		PageBefore:  before,
		PageAfter:   after,
		TimeStep:    step,
	}

}
