// run.go
package main

import (
	"fmt"
	"server"
)

func main_() {
	fmt.Println("Hello World!")
	var s = server.NewIRCServer()
	err := s.Start()

	fmt.Println(err)

	<-s.Done
}
