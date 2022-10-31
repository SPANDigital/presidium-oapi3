package markdown

import "fmt"

// GetWeightedFilename prefixes the filename with a digit
func GetWeightedFilename(weight int, filename string) string {
	var wFilename string

	// Ensure number of digits is consistent for smaller numbers
	if weight < 10 {
		wFilename = fmt.Sprintf("0%v-%v", weight, filename)
	} else {
		wFilename = fmt.Sprintf("%v-%v", weight, filename)
	}

	return wFilename
}
