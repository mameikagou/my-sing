package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file1.Close()

	// 创建文件
	file2, err := os.Create("dest.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file2.Close()

	writer := bufio.NewWriter(file2)
	reader := bufio.NewReader(file1)

	_, err = io.Copy(writer, reader)

	// err = writer.Flush() 用于将缓冲区中的数据写入到底层的 io.Writer。
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
	fmt.Println("Wrote to file using bufio.Writer")
}
