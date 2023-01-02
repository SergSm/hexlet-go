package main 

import (
	"./greeting" // Относительный импорт
	"fmt" 
)

func main() {
	// That's how it works:
	// fmt.Println(greeting.Get())	
	// output: Hello, hexlet

	// And this is how it's nt
	//fmt.Println(greeting.greeting) 
	// output:
	//./main.go:14:23: greeting not exported by package greeting

	// Aliases example:
	//fmt.Println("First greet", greetingv1.Get(), "\n", "Second greet", greetingv2.Get())

	fmt.Println(greeting.Hello())
}