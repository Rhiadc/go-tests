package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(person string) string {
	if person == "" {
		person = "Go"
	}
	return englishPrefix + person
}

func main() {
	fmt.Println(Hello("Rhiad"))
}
