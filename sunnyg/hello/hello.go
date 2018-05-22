package hello

import (
	"fmt"
)

const langEN = "en"
const langES = "es"
const langFR = "fr"
const prefixEN = "Hello, %s"
const prefixES = "Hola, %s"
const prefixFR = "Bonjour, %s"

const defaultName = "World"

// Prints a greeting in English, Spanish or French.
func Hello(name string, lang string) string {
	prefix := greetingPrefix(lang)

	if name == "" {
		name = defaultName
	}

	return fmt.Sprintf(prefix, name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case langEN: 	prefix = prefixEN
	case langES: 	prefix = prefixES
	case langFR: 	prefix = prefixFR
	default: 		prefix = prefixEN
	}
	return
}
