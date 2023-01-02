package main

import (
	"fmt"
)

const (
	Spanish = "Spanish"
	French  = "French"
	English = "English"
)

var greetingPrefix = map[string]string{
	"":      "Hello",
	English: "Hello",
	Spanish: "Hola",
	French:  "Salut",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix[language] + ", " + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
