package reader

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sysops/types"
	"sysops/validation"
)

//Reader reads input from a specified file and stores it as a string
type Reader struct {
	RawData  string
	Path     string
	CommandQ chan *types.Request
}

//NewReader initialized an empty Reader object
func NewReader() *Reader {

	return &Reader{
		Path:     "", //file path to be read
		RawData:  "", //buffer to store read input as a string
		CommandQ: make(chan *types.Request, 100),
	}
}

//Decode decodes commands
func (r *Reader) Decode() {

	if len(r.RawData) == 0 {
		println("Data Empty")
		return
	}

	commands := strings.Split(r.RawData, "\n")

	for _, command := range commands {

		str := strings.Fields(command)

		if len(str) == 0 {
			println("Empty line....skipping")
		} else {

			//validate commands
			ok, req := validation.Validate(str)
			if ok {
				r.CommandQ <- req
				// fmt.Println("Command Received: ", str)
			} else {
				fmt.Println("Invalid Command", str)
			}

		}
	}

}

//ReadFile reads a file and saves it to raw Data
func (r *Reader) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	stringData := string(data)
	r.RawData = stringData

}
