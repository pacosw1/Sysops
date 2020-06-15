package types

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
