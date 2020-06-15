package main

import (
	"fmt"
	"sysops/system"
)

func main() {

	pageSize := 16 //bytes

	manager := system.New(2048, 4096, pageSize)

	manager.Reader.ReadFile("files/test.txt")
	manager.Reader.Decode()

	manager.Start()
	// p1 := types.NewProcess(4, 48, 16)
	// p2 := types.NewProcess(3, 2048, 16)

	// manager.LoadProcess(p1)
	// manager.LoadProcess(p2)

	// fmt.Println(len(manager.Monitor.Requests))

	for i := 0; i < len(manager.Monitor.Requests); i++ {

		req := manager.Monitor.Requests[i]

		fmt.Printf("Type: " + req.Type + "\n")

		for j := 0; j < len(req.Logs); j++ {
			log := req.Logs[j]
			log.Print()
		}

	}

}
