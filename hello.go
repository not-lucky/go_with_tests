// test
package main

import (
	"fmt"
)

// Hello ...
func Hello(name string) string {
	const greet = "Hello, "
	if name == "" {
		return greet + "world"
	}
	return greet + name
}

func main() {
	fmt.Println(Hello("world"))
}
