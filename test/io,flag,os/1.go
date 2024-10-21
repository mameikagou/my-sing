package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./example.txt")
	fmt.Print(os.Getwd())
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)

	for {
		// Read(p []byte) (n int, err error)： 从输入流中读取数据到 p 中，返回读取的字节数和可能的错误。
		// 把file的内容写入buf当中
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Print(err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
