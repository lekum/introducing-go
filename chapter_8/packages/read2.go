package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil {
		return
	}
	buf := make([]byte, fInfo.Size())
	n, err := f.Read(buf)
	if err != nil {
		return
	}
	fmt.Println("Read", n, "bytes")
	fmt.Println(string(buf))
}
