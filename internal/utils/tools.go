package utils

import (
	"fmt"
	"strings"
)

func FormatBinary(n uint64, bitSize int) string {
	// Convert to binary string
	binStr := fmt.Sprintf("%0*b", bitSize, n)

	// Separate every 8 bits
	var parts []string
	for i := 0; i < len(binStr); i += 8 {
		parts = append(parts, binStr[i:i+8])
	}

	// Join with spaces
	return strings.Join(parts, " ")
}
