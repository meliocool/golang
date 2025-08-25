package main

// * Package Time consists of time management related functions

import (
	"fmt"
	"time"
)

func main() {
	// * time.Now() -> Get the current time
	now := time.Now()
	fmt.Println(now.Local())

	//* time.Date() -> Create a Date
	utc := time.Date(2005, time.August, 5, 12, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// * time.Parse(layout, string) -> Parse time from string
	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 12:04:03"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	// * Duration -> Type which is essentially just an alias for int64
	// * a LOT of method exists for Duration

	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 60
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration: %d\n", duration3)

}
