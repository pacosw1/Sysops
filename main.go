package main

import (
	"sysops/system"
	"sysops/types"
)

func main() {

	pageSize := 16 //bytes
	//manage system
	manager := system.New(350, 3000, pageSize)
	manager.ReadFile("files/test.txt")

	// list := manager.PhysicalMem.FreePages.Front()

	p := types.NewProcess(20, 242, pageSize)
	p2 := types.NewProcess(30, 120, pageSize)

	manager.LoadProcess(p)
	manager.LoadProcess(p2)
	// for list != nil {
	// 	fmt.Println(list.Value)
	// 	list = list.Next()
	// }

	manager.PhysicalMem.View()
	println("TimeStep: ", manager.TimeStep)
	// fmt.Println(p.Pages[0])

}
