package routines

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func ChannelFunc(ch chan<- string) {
	time.Sleep(5 * time.Second)
	ch <- "Hello Channel"
}
