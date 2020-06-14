package replacement

import (
	"container/list"
	"sysops/types"
)

//FIFO first in first out struct
type FIFO struct {
	OrderList *list.List
}

func NewFIFO() *FIFO {
	return &FIFO{
		OrderList: list.New(),
	}
}

//Next returns the next value to be removed
func (f *FIFO) Next() *types.Page {
	if f.OrderList.Len() == 0 {
		return nil
	} else {
		back, _ := f.OrderList.Back().Value.(*types.Page)
		return back
	}
}

//Empty checks if list is empty
func (f *FIFO) Empty() bool {
	return f.OrderList.Len() == 0
}

//adds Page to the front of the list, used when introducing program to memory
func (f *FIFO) Add(page *types.Page) {
	f.OrderList.PushFront(page)

}

//Remove returns and removes the next value in list back, LRUsed value
func (f *FIFO) Remove() *types.Page {

	last, _ := f.OrderList.Back().Value.(*types.Page)
	f.OrderList.Remove(f.OrderList.Back())
	return last
}
