package main

import "fmt"

func main() {
	s := "elite"

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v %d\n", []byte(s), len([]byte(s)))
}
