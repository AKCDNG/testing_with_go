package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const italianGreetingPrefix = "Ciao, "
const spanish = "Spanish"
const french = "French"
const italian = "Italian"

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
	case italian:
		prefix = italianGreetingPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Anish", "Spanish"))
}
