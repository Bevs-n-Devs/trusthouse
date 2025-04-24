package utils

import (
	"strings"
	"unicode"
)

// CleanPostcode takes a raw postcode and removes all whitespace and uppercases it
func CleanPostcode(rawPostcode string) string {
	var cleanPostcode strings.Builder
	cleanPostcode.Grow(len(rawPostcode))

	// Iterate over each rune in the string and add it to the cleanPostcode
	// only if it is not whitespace
	for _, r := range rawPostcode {
		if !unicode.IsSpace(r) {
			cleanPostcode.WriteRune(unicode.ToUpper(r))
		}
	}
	return cleanPostcode.String()
}
