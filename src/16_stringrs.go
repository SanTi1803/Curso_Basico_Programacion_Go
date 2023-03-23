package main

import "fmt"

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC PC) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.Ram, myPC.Disk, myPC.Brand)
}

func main() {
	myPC := PC{Ram: 16, Brand: "MSI", Disk: 500}

	fmt.Println(myPC)
}
