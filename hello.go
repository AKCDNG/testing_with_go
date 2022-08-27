package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishGreetingPrefix
	
	switch language {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Anish", "Spanish"))
}
