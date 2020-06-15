package replacement

import "sysops/types"

//Algo re
type Algo interface {
	Print()
	Push(p *types.Page)
	Remove(p *types.Page)
	Pop() *types.Page
	Peek() *types.Page
	Empty() bool
}
