package main

import (
	"fmt"
	"path"
	"path/filepath"
)

// * Package path is used to manipulate data path like in URL or File System
// * By default uses Slash '/', therefore good for URL
// * For File System, use path/filepath because windows uses backslash '\'

func main() {
	// * path
	fmt.Println(path.Dir("hello/world.go")) // * Creates a PATH not a DIRECTORY
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	// * filepath -> Check if the OS is unix based (slash) or windows (backslash)
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
