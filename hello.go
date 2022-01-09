package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(name) > 0 {
		return englishHelloPrefix + name
	}
	return englishHelloPrefix + "World"
}

func main() {
	fmt.Println(Hello("Niko"))
}
