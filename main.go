package main

import (
	"fmt"

	"github.com/ramamimu/golang-modules/hello"
)

func main() {
	fmt.Println("Hello World")
	hello.JustPrint()
	called:=hello.CallGetHello("Rama")
	fmt.Println(called)
}
