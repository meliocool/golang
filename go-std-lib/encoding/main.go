package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// * Encoding allows to encode and decode data like base64, csv, and json

func main() {
	// * Base64
	var encoded = base64.StdEncoding.EncodeToString([]byte("Hi My Name Is Lee Jae-In"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	// * CSV Reader
	csvString := "muhammad,dhitan,imam\n" + "saleh,muhammad,segeir\n" + "aldiaz,kusuma,ramadhan\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	// * CSV Writer
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"muhammad", "dhitan", "imam"})
	_ = writer.Write([]string{"saleh", "muhammad", "segeir"})
	_ = writer.Write([]string{"aldiaz", "kusuma", "ramadhan"})

	writer.Flush()
}
