package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := "Anonymous"
	if len(os.Args) < 2 {
		log.Printf("Name missing, assuming anonymous")
	} else {
		name = os.Args[1]
	}
	fmt.Println("Hello ", name)
	fmt.Println("hi")
}
