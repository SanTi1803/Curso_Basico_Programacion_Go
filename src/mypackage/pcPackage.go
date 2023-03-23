package mypackage

import "fmt"

// PC struct with public access
type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func (myComputer *PC) DuplicateRAM() {
	myComputer.Ram = myComputer.Ram * 2
}

//PUBLIC Da a conocer el "Brand" de la "Pc"
func (myPc PC) ping() {
	fmt.Println(myPc.Brand, "Pong")
}
