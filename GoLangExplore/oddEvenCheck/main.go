package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range arr {
		if v%2 == 0 {
			fmt.Printf("%v is Even\n", i)
			continue
		}
		fmt.Printf("%v is Odd\n", i)
	}
}
