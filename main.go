package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	f, err := os.Open("io_test.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		fmt.Printf("> Read %d character\n", len(line))
		fmt.Printf(">>%s\n", limitLength(line, 50))

		if err != nil {
			break
		}
	}
}

func limitLength(s string, length int) string {
	if len(s) < length {
		return s
	}
	return s[:length]
}

func main() {
	ReadFile()
}
