package replacement

import (
	"container/list"
	"fmt"
	"sysops/types"
)

//FIFO first in first out struct
//O(1) Search O(1) modifying O(N) Memory
type FIFO struct {
	OrderList *list.List               //store the page on a node of the list
	Location  map[string]*list.Element //Stores the location and node of page on list
}

//NewFIFO creates and initializes a new fifo struct
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

//Push adds Page to the front of the list, used when introducing program to memory
func (f *FIFO) Push(page *types.Page) {

	UID := parseUID(page)

	node := f.Location[UID]

	//already exists in list, do nothing
	if node != nil {
		return
	}

	//else add it to structure
	f.OrderList.PushFront(page)
	f.Location[UID] = f.OrderList.Front()
}

//Print prints queue
func (f *FIFO) Print() {

	dummy := f.OrderList.Front()

	//don't do anything if empty
	if f.Empty() {
		return
	}

	for dummy != nil {
		// fmt.Println(dummy.Value)
		val, _ := dummy.Value.(*types.Page)
		fmt.Printf("PID: %d Page#: %d  -->  ", val.PID, val.ID)
		dummy = dummy.Next()
	}
}

//Remove removes a page from replacement queue
func (f *FIFO) Remove(p *types.Page) {

	//generate unique ID
	UID := parseUID(p)
	node := f.Location[UID]

	//don't do anything if empty
	if node == nil {
		return
	}

	f.OrderList.Remove(node)
	delete(f.Location, UID)

}

//Pop returns and removes the next value in list back, LRUsed value
func (f *FIFO) Pop() *types.Page {

	//type assertion in go
	last, _ := f.OrderList.Back().Value.(*types.Page)

	//delete from both linked list and hash
	f.OrderList.Remove(f.OrderList.Back())
	delete(f.Location, parseUID(last))
	return last

}
