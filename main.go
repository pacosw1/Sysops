package main

import (
	"fmt"
	"sysops/types"
)

func main() {

	// pageSize := 16 //bytes

	mem := types.NewProcess(420, 13, 4)

	for i := 0; i < mem.Size; i++ {
		bit := mem.Memory[i]
		fmt.Printf("Virtual Address: %d  Virtual Page: %d  Offset: %d  \n\n", i, bit.Page, bit.Offset)
	}

	//manage system
	// manager := system.New(2049, 3000, pageSize)

	// manager.Reader.ReadFile("files/test.txt")
	// manager.Reader.Decode()

	// list := manager.PhysicalMem.FreePages.Front()

	// p := types.NewProcess(20, 242, pageSize)
	// p2 := types.NewProcess(30, 120, pageSize)
	// manager.LoadProcess(p)
	// manager.LoadProcess(p2)
	// for list != nil {
	// 	fmt.Println(list.Value)
	// 	list = list.Next()
	// }
	// manager.Start()

	// println("TimeStep: ", manager.TimeStep)
	// fmt.Println(p.Pages[0])

}
