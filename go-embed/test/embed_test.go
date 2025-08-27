package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// * go:embed version.txt will store the txt file into the string variable

// * Embed File to String (File must be filled with chars too)

//go:embed testdata/version.txt
var version string

//go:embed testdata/version.txt
var version2 string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

// * Embed File to []byte
// * This is fit for embedding binary files, like images or audio

//go:embed testdata/gyj5.jpg
var gyj []byte

func TestByteArray(t *testing.T) {
	err := os.WriteFile("gyj_next.jpg", gyj, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// * Embed Multiple Files
// * Change the variable type with embed.FS

//go:embed testdata/files/a.txt
//go:embed testdata/files/b.txt
//go:embed testdata/files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, Aerr := files.ReadFile("files/a.txt")
	if Aerr != nil {
		panic(Aerr)
	}
	fmt.Println(string(a))

	b, Berr := files.ReadFile("files/b.txt")
	if Berr != nil {
		panic(Berr)
	}
	fmt.Println(string(b))

	c, Cerr := files.ReadFile("files/c.txt")
	if Cerr != nil {
		panic(Cerr)
	}
	fmt.Println(string(c))
}

// * Path Matcher
// * We can use Path Matcher to read multiple files
// * Use regex

// * ALL files with the extension.txt

//go:embed ../files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
