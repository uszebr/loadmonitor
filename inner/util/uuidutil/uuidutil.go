package uuidutil

import (
	"hash/fnv"

	"github.com/google/uuid"
)

// Predefined vibrant color palette
var vibrantColors = []string{
	"#FF5733", // Red-orange
	"#FF8D1A", // Bright orange
	"#FFC300", // Yellow
	"#DAF7A6", // Light green
	"#33FF57", // Green
	"#33CFFF", // Sky blue
	"#3375FF", // Blue
	"#8D33FF", // Purple
	"#FF33A6", // Pink
}

// ColorFromUUID returns a vibrant color based on a given UUID
func ColorFromUUID(u uuid.UUID) string {
	h := fnv.New32a()
	h.Write(u[:]) // Convert UUID to a byte slice
	hashValue := h.Sum32()

	// Select a color from the predefined palette based on the hash value
	color := vibrantColors[hashValue%uint32(len(vibrantColors))]
	return color
}

// First4Symbols returns the first 4 characters of a UUID string
func First4Symbols(u uuid.UUID) string {
	uuidStr := u.String()
	if len(uuidStr) < 4 {
		return uuidStr
	}
	return uuidStr[:4]
}
