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

//NewCommentReq constructor
func NewCommentReq(input []string) *CommentReq {

	return &CommentReq{
		Type:  globals.Print,
		Input: input,
	}
}

//GenerateOutput generates response based on output and logs
func (f *CommentReq) GenerateOutput(output string) {

	fmt.Println("-------------------------------------------------------------------\n")
	fmt.Println(f.Input, "START\n")
	fmt.Println("-------------------------------------------------------------------\n\n")

	fmt.Println(f.Input[1:])

}

//Args gets input arguments
func (f *CommentReq) Args() []int {
	return []int{}
}

//GetType gets the type for this command
func (f *CommentReq) GetType() string {
	return f.Type
}

//AddLog Adds log to this commands
func (f *CommentReq) AddLog(p *PageLog) {

}
