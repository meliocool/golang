package main

// * os -> Package to access OS related information, such as Args, Hostname, Creating files, env, and so on

import (
	"fmt"
	"os"
)

func main() {
	// * Args -> Arguments after running a go file
	// * go run main.go Hi! this is an argument
	// * The argument will be stored in a Slice, use "" if want to be united in 1 line
	args := os.Args
	for _, arg := range args {
		fmt.Println("Arguments: ", arg)
	}

	// * Hostname -> The Computer HostName
	// * returns the name and an error boolean
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println(err)
	}

	// * Create & write file
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	if _, err := file.WriteString("Hello from os package!\n"); err != nil {
		file.Close()
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}

	// * Open (read) then close
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	// * ... optionally read here ...
	if err := f.Close(); err != nil {
		panic(err)
	}

	// * Safe to remove!
	if err := os.Remove("example.txt"); err != nil {
		fmt.Println("Remove error:", err)
	}

	// * ENV
	_ = os.Setenv("MY_ENV", "secret")
	fmt.Println("MY_ENV:", os.Getenv("MY_ENV"))
	if value, ok := os.LookupEnv("PATH"); ok {
		fmt.Println("PATH is set (length):", len(value)) // avoid huge print
	}

	// * Working directory
	if dir, err := os.Getwd(); err == nil {
		fmt.Println("Current Directory:", dir)
	}
	if err := os.Chdir(".."); err == nil {
		if newDir, err := os.Getwd(); err == nil {
			fmt.Println("Now in:", newDir)
		}
	}

	// * File info
	if info, err := os.Stat("go.mod"); err == nil {
		fmt.Println("File Name:", info.Name())
		fmt.Println("Size:", info.Size())
		fmt.Println("IsDir:", info.IsDir())
	} else {
		fmt.Println("Stat go.mod error:", err)
	}
}
