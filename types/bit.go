package types

//Info stores useful data based on a bit virtual Address
type Info struct {
	Offset int //offset in page for bit
	Page   int //page number for bit
}

//NewInfo intializes info struct
func NewInfo(offset, page int) *Info {
	return &Info{
		Offset: offset,
		Page:   page,
	}
}
