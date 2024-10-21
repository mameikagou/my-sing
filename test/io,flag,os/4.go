package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio 包提供了对输入和输出的缓冲功能。它通过在读取或写入时使用缓冲区来提高 I/O 操作的效率。bufio 包中的类型和函数可以帮助你更高效地处理 I/O 操作，特别是当你需要频繁读取或写入小块数据时。
func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
