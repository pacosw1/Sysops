package replacement

import (
	"container/list"
	"fmt"
	"strconv"
	"sysops/types"
)

//LRU struct
type LRU struct {
	OrderList *list.List
	Hash      map[string]*list.Element
}

//NewLRU initializes new LRU
func NewLRU() *LRU {
	return &LRU{
		OrderList: list.New(),
		Hash:      make(map[string]*list.Element, 0),
	}
}

//Empty checks if LRU is empty
func (l *LRU) Empty() bool {
	return l.OrderList.Len() == 0
}

//Peek returns the next value to be removed
func (l *LRU) Peek() *types.Page {
	back, ok := l.OrderList.Back().Value.(*types.Page)

	if !ok {
		return nil
	}
	return back
}

//Print prints the LRU queue
func (l *LRU) Print() {

	dummy := l.OrderList.Front()

	if l.Empty() {
		return
	}

	for dummy != nil {
		// fmt.Println(dummy.Value)
		val, _ := dummy.Value.(*types.Page)
		fmt.Printf("PID: %d Page#: %d  -->  ", val.PID, val.ID)
		dummy = dummy.Next()
	}
}

//Remove removes an item from the queue
func (l *LRU) Remove(p *types.Page) {

	PID := parseUID(p)

	node := l.Hash[PID]

	if node == nil {
		return
	}

	//deletes page from list
	l.OrderList.Remove(node)
	//deletes page from hash tables
	delete(l.Hash, PID)

}

//Push adds Page to the front of the list, used when introducing program to memory
func (l *LRU) Push(page *types.Page) {

	//create a unique id using processID and pageID
	UID := parseUID(page)

	//nill if not inside hash
	node := l.Hash[UID]
	//if value already in our list
	if node != nil {
		l.use(UID)
		return
	}

	//add page to the front of our list if not present
	l.OrderList.PushFront(page)
	//add node to the hash table
	l.Hash[UID] = l.OrderList.Front()

}

//Use uses a value, and thus moves it to the most recently used position in list (front)
func (l *LRU) use(UID string) {
	//create unique ID
	node := l.Hash[UID]
	l.OrderList.MoveToFront(node)
	l.Hash[UID] = node

}

//Generates a unique id made up by page number and process ID
func parseUID(p *types.Page) string {
	return strconv.Itoa(p.PID) + strconv.Itoa(p.ID)
}

//Pop returns and removes the next value in list back, LRUsed value
func (l *LRU) Pop() *types.Page {

	//type assertion
	back, ok := l.OrderList.Back().Value.(*types.Page)

	if !ok {
		fmt.Print("LRU empty")
		return nil
	}

	//deletes page from list
	l.OrderList.Remove(l.OrderList.Back())
	//deletes page from hash tables
	delete(l.Hash, parseUID(back))

	return back
}
