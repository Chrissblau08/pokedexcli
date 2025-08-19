package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// alles kleinschreiben
	text = strings.ToLower(text)

	// leading/trailing whitespace wegschneiden
	text = strings.TrimSpace(text)

	// nach whitespace splitten
	words := strings.Fields(text)

	return words
}
