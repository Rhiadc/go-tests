package main

import (
	"fmt"
)

const (
	spanish = "spanish"
	english = "english"
	french  = "french"

	spanishPrefix = "Hola, "
	englishPrefix = "Hello, "
	franchPrefix  = "Bounjour, "
)

func Hello(person, language string) string {
	if person == "" {
		person = "Go"
	}
	return switchLanguage(language) + person
}

func switchLanguage(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = franchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Rhiad", "English"))
}
