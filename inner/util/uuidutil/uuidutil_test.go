package uuidutil

import (
	"fmt"
	"hash/fnv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestColorFromUUID tests the ColorFromUUID function with unique UUID test cases
func TestColorFromUUID(t *testing.T) {
	testCases := []struct {
		uuidStr       string
		expectedColor string
	}{
		{"123e4561-e83b-12d3-a456-426621230000", vibrantColors[hashModulo(uuid.MustParse("123e4561-e83b-12d3-a456-426621230000"))]},
		{"123e3567-e89b-12d3-a456-426612342401", vibrantColors[hashModulo(uuid.MustParse("123e3567-e89b-12d3-a456-426612342401"))]},
		{"123e4547-e89b-12d3-a456-426234634002", vibrantColors[hashModulo(uuid.MustParse("123e4547-e89b-12d3-a456-426234634002"))]},
		{"123e4557-e89b-12d3-a456-426613544003", vibrantColors[hashModulo(uuid.MustParse("123e4557-e89b-12d3-a456-426613544003"))]},
	}

	for _, tc := range testCases {
		t.Run(tc.uuidStr, func(t *testing.T) {
			u, err := uuid.Parse(tc.uuidStr)
			assert.NoError(t, err, "UUID parsing failed")

			color := ColorFromUUID(u)
			assert.Equal(t, tc.expectedColor, color, fmt.Sprintf("Color for UUID %s should be %s", tc.uuidStr, tc.expectedColor))
		})
	}
}

// Helper function to compute the expected index in the vibrantColors array based on UUID
func hashModulo(u uuid.UUID) int {
	h := fnv.New32a()
	h.Write(u[:]) // Convert UUID to byte slice
	return int(h.Sum32() % uint32(len(vibrantColors)))
}

// TestFirst4Symbols tests the First4Symbols function with various test cases
func TestFirst4Symbols(t *testing.T) {
    testCases := []struct {
        name         string
        uuidStr      string
        expected     string
    }{
        {
            name:     "Standard UUID",
            uuidStr:  "123e4567-e89b-12d3-a456-426614174000",
            expected: "123e",
        },
        {
            name:     "Different UUID",
            uuidStr:  "987f6543-21cb-98de-b123-543210abcd00",
            expected: "987f",
        },
        {
            name:     "Another UUID",
            uuidStr:  "abcdefab-cdef-abcd-efab-cdefabcdef01",
            expected: "abcd",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            u, err := uuid.Parse(tc.uuidStr)
            assert.NoError(t, err, "UUID parsing failed")

            result := First4Symbols(u)
            assert.Equal(t, tc.expected, result, "Expected first 4 characters of UUID to be %s but got %s", tc.expected, result)
        })
    }
}
