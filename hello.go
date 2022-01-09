package main

import "fmt"

const defaultName = "World"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if len(name) == 0 {
		return englishHelloPrefix + defaultName
	}

	return getPrefix(lang) + name
}

// (prefix string) = named return type
func getPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Niko", ""))
}
