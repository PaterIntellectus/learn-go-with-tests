package helloworld

type language = string

const (
	english language = "en"
	russian language = "ru"
	spanish language = "es"
	french  language = "fr"

	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	helloSuffix = "!"
)

func Hello(name string, language language) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + helloSuffix
}

func greetingPrefix(language language) (prefix string) {
	switch language {
	default:
	case english:
		prefix = englishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return
}
