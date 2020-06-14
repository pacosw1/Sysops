package main

import (
	"sysops/globals"
	"sysops/system"
	"sysops/types"
)

func main() {

	pageSize := 4 //bytes

	// mem := types.NewProcess(420, 9, 4)

	// println(len(mem.Pages))

	// for i := 0; i < mem.Size; i++ {
	// 	bit := mem.Memory[i]
	// 	fmt.Printf("Virtual Address: %d  Virtual Page: %d  Offset: %d  \n\n", i, bit.Page, bit.Offset)
	// }

	// //manage system

	manager := system.New(globals.MaxSize, 100, pageSize)

	// p1 := types.NewProcess(1, 32, pageSize)
	p2 := types.NewProcess(2, 5, pageSize)
	p3 := types.NewProcess(3, 50, pageSize)

	// // // fmt.Println(len(p1.Pages))
	// manager.LoadProcess(p1)
	// manager.LoadProcess(p2)
	manager.LoadProcess(p3)
	manager.LoadProcess(p2)
	manager.FreePages(p3)

	manager.SwapMemory.View()

	// for _, page := range p3.Pages {

	// 	fmt.Printf("PID: %d ID: %d pageFrame: %d swampAddr: %d \n ", page.PID, page.ID, page.PageFrame, page.SwapAddr)
	// }
	// manager.FreePages(p1)
	// manager.LoadProcess(p2)
	// manager.LoadProcess(p3)

	// manager.PhysicalMem.View()

	// lru := replacement.NewLRU()

	// p1 := types.NewPage(1, 420)
	// p2 := types.NewPage(2, 69)
	// p3 := types.NewPage(3, 52)

	// lru.Push(p1)
	// lru.Push(p2)
	// lru.Push(p3)

	// lru.Push(p1)

	// lru.Pop()

	// lru.Print()

	// println(len(lru.Hash))

	//debe de quedar  p1 p2 p3
	// 						^ mas viejo

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
