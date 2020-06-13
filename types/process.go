package types

//Process struct
type Process struct {
	pID      int
	Size     int           //size of process in bytes
	Pages    map[int]*Page //table for pages in process
	PageSize int
}

//NewProcess create new Process
func NewProcess(id, size, pageSize int) *Process {
	p := &Process{
		pID:      id,
		Size:     size,
		PageSize: pageSize,
		Pages:    make(map[int]*Page, 0),
	}
	p.Init()
	return p
}

//Init Process Page Table
func (p *Process) Init() {
	pages := p.Size / p.PageSize
	//handle case for leftover values
	if p.Size%p.PageSize != 0 {
		pages++
	}
	for i := 0; i < pages; i++ {
		p.Pages[i] = NewPage(p.pID, i)
	}
}
