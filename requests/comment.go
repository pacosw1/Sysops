package requests

import (
	"fmt"
	"sysops/globals"
)

//CommentReq request to free mem
type CommentReq struct {
	Type  string
	Input []string
	PID   int
	Logs  []*PageLog
}

//NewCommentReq n
func NewCommentReq(input []string) *CommentReq {

	return &CommentReq{
		Type:  globals.Print,
		Input: input,
	}
}

//GenerateOutput p
func (f *CommentReq) GenerateOutput() {

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	fmt.Println(f.Input[0])
	fmt.Println(f.Input[1:])

	fmt.Println("-------------------------------------------------------------------\n\n")

}

//Args a
func (f *CommentReq) Args() []int {
	return []int{}
}

//Type t
func (f *CommentReq) GetType() string {
	return f.Type
}

//Type t
func (f *CommentReq) AddLog(p *PageLog) {

}
