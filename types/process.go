package types

//Process struct
type Process struct {
	PID      int           //Process ID number
	Size     int           //size of process in bytes
	Pages    map[int]*Page //table for pages in process
	Memory   []*Info       //Virtual Memory for process
	PageSize int           //Page size
}

//NewProcess initializes new Process object
func NewProcess(id, size, pageSize int) *Process {
	p := &Process{
		PID:      id,
		Size:     size,
		PageSize: pageSize,
		Pages:    make(map[int]*Page, 0),
		Memory:   make([]*Info, size),
	}
	p.Init() //run unit before returning
	return p
}

//GetInfo gets Virtual page and offset for specified virtual address
func (p *Process) GetInfo(vAddr int) *Info {
	return p.Memory[vAddr]
}

//Init initializes virtual memory and Page Table
func (p *Process) Init() {
	pagNum := 0

	//init page 0
	p.Pages[0] = NewPage(p.PID, pagNum)

	for bit := 0; bit < p.Size; bit++ {

		//calculates offset based on virtual adress and page size
		offset := bit % p.PageSize
		//save Info struct for each bit P, D to speed up calculations
		p.Memory[bit] = &Info{
			Offset: offset,
			Page:   pagNum,
		}
		//increase page size
		if ((bit)%p.PageSize) == 0 && bit != 0 {
			pagNum++
			p.Pages[pagNum] = NewPage(p.PID, pagNum) //initialized new page

		}
	}

}
