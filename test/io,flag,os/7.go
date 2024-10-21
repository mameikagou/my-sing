package main

import (
	"flag"
	"fmt"
)

var numberFlag = flag.Int("n", 0, "an int") // go run 7.go -n 42

func main() {
	// 解析命令行标志
	flag.Parse()

	// 使用标志的值
	fmt.Println("Value of n:", *numberFlag)
}
