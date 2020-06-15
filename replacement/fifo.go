package replacement

import (
	"container/list"
	"fmt"
	"sysops/types"
)

//FIFO first in first out struct
type FIFO struct {
	OrderList *list.List
	Location  map[string]*list.Element
}

func NewFIFO() *FIFO {
	return &FIFO{
		OrderList: list.New(),
		Location:  make(map[string]*list.Element, 0),
	}
}

//Peek returns the next value to be removed
func (f *FIFO) Peek() *types.Page {
	if f.OrderList.Len() == 0 {
		return nil
	}

	back, _ := f.OrderList.Back().Value.(*types.Page)
	return back

}

//Empty checks if list is empty
func (f *FIFO) Empty() bool {
	return f.OrderList.Len() == 0
}

//Add adds Page to the front of the list, used when introducing program to memory
func (f *FIFO) Add(page *types.Page) {
	f.OrderList.PushFront(page)
	f.Location[parseUID(page)] = f.OrderList.Front()
}

func (f *FIFO) Remove(p *types.Page) {

	UID := parseUID(p)
	node := f.Location[UID]

	if node == nil {
		fmt.Println("Page not present on list")
		return
	}

	f.OrderList.Remove(node)
	delete(f.Location, UID)

}

//Remove returns and removes the next value in list back, LRUsed value
func (f *FIFO) Pop() *types.Page {

	last, _ := f.OrderList.Back().Value.(*types.Page)
	f.OrderList.Remove(f.OrderList.Back())

	delete(f.Location, parseUID(last))
	return last

}
