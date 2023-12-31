package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort("tcp", "localhost", 1314)
	fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)

}
