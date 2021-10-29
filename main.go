package main

import (
	"fmt"
	"os"
)

// Create a program which lists environment variables as '=' delimited key/value pairs:
func main() {
	// Environ returns a copy of strings([]string) representing the environment,
	// in the form "key=value".
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
