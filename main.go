package main

import "fmt"

func Hello(person string) string {
	return "Hello, " + person
}

func main() {
	fmt.Println(Hello("Rhiad"))
}
