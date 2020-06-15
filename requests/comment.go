package requests

import (
	"sysops/globals"
	"sysops/monitor"
)

//FreeMem request to free mem
type CommentReq struct {
	Type  string
	Input []string
	PID   int
	Logs  []*monitor.PageLog
}

//NewFreeMemReq n
func NewCommentReq(input []string) *CommentReq {

	return &CommentReq{
		Type:  globals.Print,
		Input: input,
	}
}

//GenerateOutput p
func (f *CommentReq) GenerateOutput() {

}

//Args a
func (f *CommentReq) Args() []int {
	return []int{}
}

//Type t
func (f *CommentReq) GetType() string {
	return f.Type
}
