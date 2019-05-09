package main

import (
	"fmt"
	"github.com/robteix/testmod"
	testmod2 "github.com/rselbach/testmod/v2"
)

func main() {
	fmt.Println("test")
	fmt.Println(testmod.Hi("roberto"))
	s, e := testmod2.Hi("test", "cool")
	fmt.Println(s,e)
}
