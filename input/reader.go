package input

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sysops/globals"
)

type Request struct {
	Type    string
	Args    []int
	Message []string
}

type InputReader struct {
	RawData  string
	Path     string
	CommandQ chan *Request
}

func NewReader() *InputReader {

	return &InputReader{
		Path:     "",
		RawData:  "",
		CommandQ: make(chan *Request, 100),
	}
}

//Decode decodes commands
func (r *InputReader) Decode() {

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
			ok, req := r.Validate(str)
			if ok {
				r.CommandQ <- req
				// fmt.Println("Command Received: ", str)
			} else {
				fmt.Println("Invalid Command", str)
			}

		}
	}

}

func (r *InputReader) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	stringData := string(data)
	r.RawData = stringData

}

//type validation
func isNumber(arg string) (bool, int64) {

	if num, err := strconv.ParseInt(arg, 10, 64); err == nil {
		return true, num
	}
	return false, -1

}

//Validate commands
func (r *InputReader) Validate(command []string) (bool, *Request) {
	action := command[0]
	switch action {
	case globals.Print:
		req := &Request{
			Type:    "C",
			Message: command[1:],
		}
		return true, req
	case globals.LoadP: //todo
		//check for invalid commands
		if len(command) > 3 || len(command) < 3 {
			return false, nil
		}
		//check that args are valid numbers
		valid, size := isNumber(command[1])
		validID, pID := isNumber(command[2])

		if valid && validID {
			if size <= globals.MaxSize {
				req := &Request{
					Type: "P",
					Args: []int{int(size), int(pID)},
				}
				return true, req
			}
		}
		return false, nil
	case globals.Access:
		//invalid command params
		if len(command) < 4 || len(command) > 4 {
			return false, nil
		}

		//check valid numbers
		validAddr, addr := isNumber(command[1])
		validProc, proc := isNumber(command[2])
		validMod, val := isNumber(command[3])

		if validAddr && validProc && validMod {
			//valid boolean value
			if val > 1 || val < 0 {
				return false, nil
			}
			req := &Request{
				Type: "A",
				Args: []int{int(addr), int(proc), int(val)},
			}
			return true, req
		}
		return false, nil

	case globals.FreeP:
		if len(command) > 2 {
			return false, nil
		}
		valid, num := isNumber((command[1]))
		if !valid {
			return false, nil
		}
		req := &Request{
			Type: "L",
			Args: []int{int(num)},
		}
		return true, req
	case globals.Stats:
		if len(command) > 1 {
			return false, nil
		}
		return true, &Request{}
	case globals.End:
		if len(command) > 1 {
			return false, nil
		}
		req := &Request{
			Type: "E",
		}
		return true, req
	default:
		fmt.Println("Command not recognized")
		return false, nil
	}
}
