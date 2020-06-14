package replacement

import "container/list"

type LRU struct {
	OrderList list.List
}


func 