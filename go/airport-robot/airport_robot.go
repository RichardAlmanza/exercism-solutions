package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct {}

func (italian Italian) LanguageName() string {
	return "Italian"
}

func (italian Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct {}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (portuguese Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf(
		"I can speak %s: %s",
		greeter.LanguageName(),
		greeter.Greet(visitorName),
	)
}


