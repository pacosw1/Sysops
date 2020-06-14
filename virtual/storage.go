package virtual

import (
	"container/list"
	"fmt"
	"sysops/types"
)

//Storage simulates real space in computer
type Storage struct {
	Memory    map[int]*types.Page //memory with physical pages
	Size      int                 //size of the structure
	PageSize  int                 //page size for memory
	FreePages *list.List          //list with free page numbers
}

//NewStorage initializes memory to specified size
func NewStorage(size, pagesize int) *Storage {
	s := &Storage{
		PageSize:  pagesize,
		Memory:    make(map[int]*types.Page, 0),
		Size:      size,
		FreePages: list.New(),
	}
	s.init()
	return s
}

//View display the current memory allocation
func (s *Storage) View() {
	mem := s.Memory

	for i := 0; i < len(mem); i++ {
		page := mem[i]
		if page == nil {
			fmt.Printf("Frame %d : <--- free \n", i)
		} else {
			fmt.Printf("Frame %d : <- PID: %d Page# %d \n", i, page.PID, page.ID)
		}
	}
}

func (s *Storage) init() {
	//determine frames in memory
	pages := s.Size / s.PageSize

	//add one extra frame if size cant be fit evenly
	if s.Size%s.PageSize != 0 {
		pages++
	}

	//all frames are free
	for page := 0; page < pages; page++ {
		s.Memory[page] = nil
		s.FreePages.PushBack(page)
	}
}
