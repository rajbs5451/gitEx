package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func main() {
	name := "Anonymous"
	if len(os.Args) < 2 {
		log.Printf(Red + "Name missing, assuming anonymous" + Reset)
	} else {
		name = os.Args[1]
	}
	fmt.Println(Green + "Hello " + name + Reset)
}
