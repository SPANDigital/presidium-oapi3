package markdown

import "fmt"

// GetWeightedFilename prefixes the filename with its weight, zero-padded so
// lexicographic file ordering matches numeric weight order (weights are now
// per-operation, so three digits are common on larger specs).
func GetWeightedFilename(weight int, filename string) string {
	return fmt.Sprintf("%04d-%v", weight, filename)
}
