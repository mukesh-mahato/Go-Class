package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println("yeah buddy!")
	fmt.Println(hello.Say(os.Args[1:]))
}
