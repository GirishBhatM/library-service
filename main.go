package main

import (
	"fmt"
	"library-service/controller"
)

func main() {
	setUp()
}

func setUp() {
	fmt.Println("Library service API V1")
	controller.MapRoutes()
}
