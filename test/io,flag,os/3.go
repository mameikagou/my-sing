package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// go run 3.go file1.txt file2.txt
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		// flag.Arg(i)
		// flag.Arg(i) 返回第 i 个非标志参数（即那些不以 - 或 -- 开头的参数）。这些参数是在调用 flag.Parse() 之后解析的。
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
