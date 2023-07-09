package main

import (
	"fmt"

	"github.com/ramamimu/golang-modules/hello"
	unittest "github.com/ramamimu/golang/unit-test"
)

func main() {
	fmt.Println("Hello World")
	hello.JustPrint()
	called:=hello.CallGetHello("Rama")
	fmt.Println(called)
	fmt.Println(hello.CallGetHello("Mimu"))
	fmt.Println(unittest.HelloUnitTest())
}
