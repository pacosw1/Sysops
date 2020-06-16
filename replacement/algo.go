package replacement

import "sysops/types"

//Algo interface requirements to implement the replacement type used in simulation
type Algo interface {
	Print()
	Push(p *types.Page)
	Remove(p *types.Page)
	Pop() *types.Page
	Peek() *types.Page
	Empty() bool
}
