package main

import "fmt"

// Map -> Data stored in Pairwise (key - value) where key is unique
// Map size is NOT FIXED, can be added as much as possible

func main() {
	person := map[string]string{
		"name":        "Lee Jae-In",
		"nationality": "South Korea",
	}
	var human map[string]string = map[string]string{} // declares an empty map
	human["name"] = "Go Youn-Jung"
	human["nationality"] = "South Korea"

	fmt.Println(person["name"], person["nationality"])
	fmt.Println(human["name"], human["nationality"])
	fmt.Println(person)
	fmt.Println(human)

	// Functions in Map
	fmt.Println("Length of Map Person: ", len(person))
	song := make(map[string]string)
	song["title"] = "Selfless"
	song["artist"] = "The Strokes"
	song["year"] = "2013"
	fmt.Println("Map Song: ", song)
	delete(song, "year")
	fmt.Println("Map song after Delete: ", song)
}
