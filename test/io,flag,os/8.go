package main

import (
	"flag"
	"fmt"
)

// go run 8.go -n

func main() {
	var numberFlag = flag.Bool("n", false, "number each line")

	flag.Parse()

	if *numberFlag {
		fmt.Println("NumberFlag mode enabled")
	} else {
		fmt.Println("NumberFlag mode disabled")
	}
}
