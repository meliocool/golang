package main

// * Package math allows for math operations, such as max, min, ceil, floor, round, asin, acos, atan, etc

import (
	"fmt"
	"math"
)

func main() {
	float1 := 6.9
	float2 := 4.2
	// * math.Round(float64) -> Round float64 up or down depending on the closest
	fmt.Println(math.Round(float1))

	// * math.Floor(float64) -> Round float64 DOWN
	fmt.Println(math.Floor(float2))

	// * math.Ceil(float64) -> Round float64 UP
	fmt.Println(math.Ceil(float1))

	// * math.Max(float64, float64) -> Returns the biggest float64 value
	fmt.Println(math.Max(float1, float2))

	// * math.Min(float64, float64) -> Returns the smallest float64 value
	fmt.Println(math.Min(float1, float2))
}
