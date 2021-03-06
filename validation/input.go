package validation

import (
	"fmt"
	"strconv"
	"sysops/globals"
	"sysops/requests"
)

//checks if input command argument is a number
func isNumber(arg string) (bool, int64) {

	if num, err := strconv.ParseInt(arg, 10, 64); err == nil {
		return true, num
	}
	return false, -1

}

//Validate commands, returns bool, and request if sucessful, if not, request will be empty
func Validate(command []string) (bool, requests.Request) {

	var req requests.Request
	//Command Type is the first argument
	action := command[0]

	//check what type of command will be validated
	switch action {

	case globals.Print: //Print, prints a message to console

		//create command
		req = requests.NewCommentReq(command)
		//return valid and request
		return true, req

	case globals.LoadP: //checks if Load Process command is valid
		//check for invalid command arguments len
		if len(command) > 3 || len(command) < 3 {
			return false, nil
		}
		//check that args are valid numbers
		valid, size := isNumber(command[1])
		validID, pID := isNumber(command[2])

		if valid && validID {
			//checks that given size is within valid ranges, max size here is declared in globals folder
			if size <= globals.MaxSize { //if all cases passed, create request and return true
				req = requests.NewLoadP(command, int(pID), int(size))
				return true, req
			}
		}
		//if not return invalid
		return false, nil

	case globals.Access: //validate Access to virtual memory command
		//check for invalid command params len
		if len(command) < 4 || len(command) > 4 {
			return false, nil
		}

		//check valid numbers
		validAddr, addr := isNumber(command[1])
		validProc, proc := isNumber(command[2])
		validMod, val := isNumber(command[3])

		//if args are valid
		if validAddr && validProc && validMod {
			//valid boolean value
			if val > 1 || val < 0 { //check bool arg validity
				return false, nil
			}
			//generate request
			req = requests.NewAccessReq(command, int(proc), int(addr), int(val))
			return true, req
		}
		return false, nil

	case globals.FreeP:
		//check for invalid command len
		if len(command) > 2 {
			return false, nil
		}
		//check if args are valid nums
		valid, num := isNumber((command[1]))
		if !valid {
			return false, nil
		}
		//generate request
		req = requests.NewFreeMemReq(command, int(num))
		return true, req
	case globals.Stats:
		//command len valid?
		if len(command) > 1 {
			return false, nil
		}
		return true, requests.NewStatsReq(command)
	case globals.End:
		//command len valid?
		if len(command) > 1 {
			return false, nil
		}
		return true, requests.NewEndReq(command)
	default:
		fmt.Println("Command not recognized")
		return false, nil
	}
}
