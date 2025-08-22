package main

import (
	"fmt"
	"sort"
)

// * Package sort is filled with sorting utilities
// * Must implement the sort.interface contract
// * Which includes the following Methods: Len, Less, Swap

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Dhitan", 20},
		{"Melio", 40},
		{"Nisha", 30},
		{"Miracle", 10},
	}
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
