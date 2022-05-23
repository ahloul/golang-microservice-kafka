package helpers

import "strings"

func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func After(value string, a string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	return value[posFirstAdjusted:]
}

func Before(value string, a string) string {
	// Get substring between two strings.
	posLast := strings.Index(value, a)
	if posLast == -1 {
		return ""
	}
	return value[:posLast]
}
