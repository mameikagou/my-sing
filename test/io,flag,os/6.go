package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	srcFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dstFile.Close()

	// 读取源文件的内容并写入到目标文件
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	_, err = io.Copy(writer, reader)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer", err)
		return
	}
	fmt.Println("Wrote to file using io.Copy")
}
