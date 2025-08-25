package main

// * flag -> Package to parse command line arguments

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	var port *int = flag.Int("port", 0, "Database port")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)
}
