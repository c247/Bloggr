package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[0:]
	if len(args) <= 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	f, err := os.Open(string(args[1]))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
