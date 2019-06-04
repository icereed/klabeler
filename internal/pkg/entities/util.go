package entities

import (
	"strings"
)

func checkIfJSONIsArray(json string) bool {
	parsedString := strings.TrimSpace(json)
	return strings.HasPrefix(parsedString, "[") && strings.HasSuffix(parsedString, "]")
}
