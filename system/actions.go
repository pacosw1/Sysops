package system

import (
	"fmt"
	"sysops/globals"
)

func validSize(size int) bool {

	return size <= globals.MaxSize

}

//FreePages loads a proces unto memory

//PrintMessage prints message sent by input
func PrintMessage(msg []string, action string) {
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
