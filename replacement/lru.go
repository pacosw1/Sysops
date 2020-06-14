package replacement

import (
	"container/list"
)

//LRU struct
type LRU struct {
	OrderList *list.List
}

//NewLRU initializes new LRU
// func NewLRU() *LRU {
// 	return &LRU{
// 		OrderList: list.New(),
// 	}
// }

// //Next returns the next value to be removed
// func (l *LRU) Next() *types.Page {

// }

// //Add adds Page to the front of the list, used when introducing program to memory
// func (l *LRU) Add(page *types.Page) {

// }

// //Use uses a value, and thus moves it to the most recently used position in list (front)
// func (l *LRU) Use(page *types.Page) {

// }

// //Remove returns and removes the next value in list back, LRUsed value
// func (l *LRU) Remove() *types.Page {

// }
