package idx

import (
	"errors"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

type incorrectCodeHandler = func(code string) (string, error)

var (
	errNotGameTDBIncorrectCode error                  = errors.New("Not GameTDB Code Error")
	incorrectCodeHandlers      []incorrectCodeHandler = []incorrectCodeHandler{
		gametdbIncorrectBCP47CodeHandler,
	}
)

const (
	incorrectGameTDBLanguageTagLen int    = 4
	subtagSeparator                string = "-"
)

func convertLanguageCodeToName(code string) string {
	name, err := parseCodeToName(code)
	if err == nil {
		return name
	}

	if _, ok := err.(language.ValueError); ok {
		return code // if code format is proper but code itself does not exist then just use the code itself, without checking further handlers
	}

	for _, handler := range incorrectCodeHandlers {
		newName, err := handler(code)
		if err == nil {
			name = newName
		}
	}

	if name == "" {
		return code // could not find meaningful name, might as well return code
	}

	return name
}

func parseCodeToName(code string) (string, error) {
	tag, err := language.Parse(code)
	if err != nil {
		return "", err
	}

	return display.Self.Name(tag), nil
}

func gametdbIncorrectBCP47CodeHandler(code string) (string, error) {
	// GameTDB uses incorrect BCP47 tag without dash
	// Since GameTDB faq acknowledges only ZHxx exceptions, the check for 4 runes is exhaustive (for now)
	if len(code) != incorrectGameTDBLanguageTagLen {
		return "", errNotGameTDBIncorrectCode
	}

	correctedCode := convertToBCP47Code(code)
	name, err := parseCodeToName(correctedCode)
	return name, err
}

func convertToBCP47Code(code string) string {
	base := strings.ToLower(string(code[0:2]))
	extension := string(code[2:4])

	return strings.Join([]string{base, extension}, subtagSeparator)
}
