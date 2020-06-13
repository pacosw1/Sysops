package types

//Page stores process block
type Page struct {
	PID        int  //process ID
	ID         int  //page ID
	InMem      bool // page is located in memory?
	SwapAddr   int  // address where page is stored in swapMem
	PageFrame  int  // page frame addr in real memory
	Mod        bool
	InsertedAt int //frame in which page was inserted to real memory
}

//NewPage creates a new blank page
func NewPage(pID, id int) *Page {
	return &Page{
		PID:        pID,
		ID:         id,
		InMem:      false,
		SwapAddr:   -1,
		PageFrame:  -1,
		Mod:        false,
		InsertedAt: -1,
	}
}
