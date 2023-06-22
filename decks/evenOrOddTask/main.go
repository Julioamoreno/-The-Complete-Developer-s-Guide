package main

import (
	"fmt"
	"math"
)

type numbers []int64

func main() {
	n := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range n {
		fmt.Println(i, "is", returnEvenOrOdd(i))
	}

}

func returnEvenOrOdd(n int) string {
	if math.Mod(float64(n), 2) == 0 {
		return "even"
	}

	return "odd"
}
