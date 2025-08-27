package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

// * When compiled into go build, the file is SAVED inside the variable
// * So if go-embed.exe is run, it will read the original file, even if the original file read is changed
// * Simple Example: version.txt is 1.0.0-SNAPSHOT, read inside the version variable, and the main is built into .exe
// * Then the version.txt is changed into 2.0.0-SNAPSHOT, it will still print out 1.0.0-SNAPSHOT
// * It is LITERALLY embedded into the variable

//go:embed test/testdata/version.txt
var version string

//go:embed test/testdata/gyj5.jpg
var gyj []byte

//go:embed test/testdata/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("gyj_next.jpg", gyj, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("test/testdata/files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("test/testdata/files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
