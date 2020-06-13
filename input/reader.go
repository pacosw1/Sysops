package input


package system

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)







type Request struct {
	Type string
	Args []int
	Message string
}







type InputReader struct {
	RawData string
	Path string
	commandQ chan *Request
}


func NewReader() *InputReader {

	return &InputReader{
		Path: ""
		RawData: "",
		commandQ: make(chan *Request, 100),
	}
}

//Decode decodes commands 
func  (r *InputReader) Decode(path string) {

	if (len(r.RawData) == 0) {
		println("Data Empty");
		return
	}

	commands := strings.Split(r.RawData, "\n")

	for _, command := range commands {

		str := strings.Fields(command)

		if len(str) == 0 {
			println("Empty line....skipping")
		} else {

			//validate commands
			println(m.Validate(str))

		}
	}

}


func  (r *InputReader) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	stringData := string(data)
	r.RawData = stringData;

}













//ReadFile reads text file with specified path
func (m *MemoryManager) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	stringData := string(data)

	commands := strings.Split(stringData, "\n")

	for _, command := range commands {

		str := strings.Fields(command)

		if len(str) == 0 {
			println("Empty line....skipping")
		} else {

			println(m.Validate(str))

		}
	}

}

//type validation
func isNumber(arg string) (bool, int64) {

	if num, err := strconv.ParseInt(arg, 10, 64); err == nil {
		return true, num
	}
	return false, -1

}

//Validate commands
func (m *MemoryManager) Validate(command []string) bool, *Request  {
	action := command[0]
	switch action {
	case Print:
		req := &Request{
			Type: "C",
			Message: command[1], 
		}
		return true, req
	case LoadP: //todo
		//check for invalid commands
		if len(command) > 3 || len(command) < 3 {
			return false
		}
		//check that args are valid numbers
		valid, size := isNumber(command[1])
		validID, pID := isNumber(command[2])

		if valid && validID {
			if size <= maxSize {
				req := &Request{
					Type: "P",
					Args: [int(size), int(pID)],
				}
				return true, req
			}
		}
		return false
	case Access:
		//invalid command params
		if len(command) < 4 || len(command) > 4 {
			return false
		}

		//check valid numbers
		validAddr, addr := isNumber(command[1])
		validProc, proc := isNumber(command[2])
		validMod, val := isNumber(command[3])

		if validAddr && validProc && validMod {
			//valid boolean value
			if val > 1 || val < 0 {
				return false
			}
			req := &Request{
				Type: "A",
				Args: [addr, proc, val],
			}
			return true, req
		}
		return false

	case FreeP:
		if len(command) > 2 {
			return false
		}
		valid, num := isNumber((command[1]))
		if !valid {
			return false
		}
		req := &Request{
			Type: "L",
			Args: [int(num)],
		}
		return true, req
	case Stats:
		if len(command) > 1 {
			return false
		}
		return true, &Request{}
	case End:
		if len(command) > 1 {
			return false
		}
		req := &Request{
			Type: "E",
			Args: [],
		}
		return true, req
	default:
		fmt.Println("Command not recognized")
		return false, &Request{}
	}
}

func printMessage(msg []string, action string) {
	res := ""
	for _, word := range msg {
		res += word + " "
	}

	fmt.Println("INPUT: " + action)

	if len(res) == 0 {
		println("OUTPUT: Empty message \n\n")
	} else {
		fmt.Print("OUTPUT: " + res + "\n\n")
		fmt.Println("")
	}
}
