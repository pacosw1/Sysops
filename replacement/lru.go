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
func (l *LRU) Peek() (*types.Page, bool) {
	back, ok := l.OrderList.Back().Value.(*types.Page)
	return back, ok
}

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

//Push adds Page to the front of the list, used when introducing program to memory
func (l *LRU) Push(page *types.Page) {

	UID := parseUID(page)

	//nill if not inside hash
	node := l.Hash[UID]
	//if value not in our list
	if node == nil {
		//store node inside hash

		// fmt.Println(newNode.Value)
		//add page to the front of our list
		l.OrderList.PushFront(page)
		//add node to the hash table
		l.Hash[UID] = l.OrderList.Front()
		// fmt.Println(l.OrderList.Front().Value)
	} else { //else use it and update its position
		l.Use(UID)
	}

}

//Use uses a value, and thus moves it to the most recently used position in list (front)
func (l *LRU) Use(UID string) {
	//create unique ID
	node := l.Hash[UID]
	l.OrderList.MoveToFront(node)
	l.Hash[UID] = node

}

func parseUID(p *types.Page) string {
	return strconv.Itoa(p.PID) + strconv.Itoa(p.ID)
}

//Pop returns and removes the next value in list back, LRUsed value
func (l *LRU) Pop() (*types.Page, bool) {

	//type assertion
	back, ok := l.OrderList.Back().Value.(*types.Page)

	if !ok {
		fmt.Print("LRU empty")
		return nil, true
	}

	//deletes page from list
	l.OrderList.Remove(l.OrderList.Back())
	//deletes page from hash tables
	delete(l.Hash, parseUID(back))

	return back, ok
}
