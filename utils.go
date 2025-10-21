package main

import "unicode"

func isLatinOnly(s string) bool {
	for _, r := range s {
		if r == ' ' || r == '-' { // Пропускает пробелы и тире.
			continue
		}

		if !(unicode.Is(unicode.Latin, r)) { // Ищет не латиницу.
			return false
		}
	}
	return true
}
