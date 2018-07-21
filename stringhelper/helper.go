package stringhelper

import (
	"fmt"
)

// Upper
func Upper(text string) string {
	return text
}

// Concat
func Concat(inputs ...string) string {
	var result string
	for _, input := range inputs {
		result = fmt.Sprintf("%s %s", result, input)
	}
	return result
}
