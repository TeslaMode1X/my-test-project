package tester

import (
	"errors"
	"fmt"
)

var username string
var language string

func Hi(name, lang string) (string, error) {
	username = name
	language = lang
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}

func FullInfoAboutUser() (string, error) {
	if username == "" || language == "" {
		return "", errors.New("user information is not set")
	}
	return fmt.Sprintf("Your name: %s \nYour age: TOP SECRET \nYour language: %s", username, language), nil
}
