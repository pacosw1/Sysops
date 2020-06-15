package types

import "fmt"

//Page stores process block
type Page struct {
	PID       int  //process ID
	ID        int  //page ID
	SwapFrame int  // address where page is stored in swapMem
	PageFrame int  // page frame addr in real memory
	Mod       bool //if memory has been modified
}

//NewPage creates a new blank page
func NewPage(pID, id int) *Page {
	return &Page{
		PID:       pID,
		ID:        id,
		SwapFrame: -1,
		PageFrame: -1,
		Mod:       false,
	}
}

//Print prints Page
func (p *Page) Print() {
	loc := ""
	frame := 1

	if p.PageFrame >= 0 {
		loc = "In Physical Memory @ Frame: "
		frame = p.PageFrame
	} else if p.SwapFrame >= 0 {
		loc = "In Secondar Storage (Swap) @ Frame: "
		frame = p.SwapFrame
	} else {
		loc = "Unassigned"
	}
	fmt.Printf("\nProcessID: %d\nFrameID: %d \n", p.PID, p.ID)
	fmt.Println(loc, frame)

}

//CopyPage deep copies the input page
func CopyPage(p *Page) *Page {
	copy := &Page{
		PID:       p.PID,
		ID:        p.ID,
		SwapFrame: p.SwapFrame,
		PageFrame: p.PageFrame,
		Mod:       p.Mod,
	}

	return copy
}
