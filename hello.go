package main

import "fmt"

const englishGreetingPrefix string = "Hello, "
const spanishGreetingPrefix string = "Hola, "
const spanish = "Spanish"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishGreetingPrefix + name + "!"
	}

	return englishGreetingPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Anish", "Spanish"))
}
