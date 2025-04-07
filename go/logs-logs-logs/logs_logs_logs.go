package logs

import (
	"unicode/utf8"
)

type AppIdentifiers struct {
	recommendation rune
	search         rune
	weather        rune
}

func NewAppIdentifiers() *AppIdentifiers {
	recommendation, _ := utf8.DecodeRuneInString("\U00002757")
	search, _ := utf8.DecodeRuneInString("\U0001F50D")
	weather, _ := utf8.DecodeRuneInString("\U00002600")
	return &AppIdentifiers{
		recommendation: recommendation,
		search:         search,
		weather:        weather,
	}
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	var r rune

	for _, char := range log {
		size := utf8.RuneLen(char)
		if size > 2 {
			r = char
			break
		}
	}

	identifiers := NewAppIdentifiers()
	if r == identifiers.recommendation {
		return "recommendation"
	} else if r == identifiers.search {
		return "search"
	} else if r == identifiers.weather {
		return "weather"
	} else {
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	modifiedLog := []rune(log)

	runeIndex := 0
	for _, char := range log {
		if char == oldRune {
			modifiedLog[runeIndex] = newRune
		}
		runeIndex++
	}

	return string(modifiedLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
