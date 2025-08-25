package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// * Bufio or Buffered IO, is used to create IO data like Reader and Writer

func main() {
	input := strings.NewReader("This is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Sup world\n")
	_, _ = writer.WriteString("Go ftw\n")
	writer.Flush()
}
