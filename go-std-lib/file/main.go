package main

import (
	"bufio"
	"io"
	"os"
)

// * File Manipulation stuff

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)

	return nil
}

func main() {
	//createNewFile("example.log", "Sample File!")

	//result, err := readFile("example.log")
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(result)
	//}

	addToFile("example.log", "\nThis is a new message!")
}
