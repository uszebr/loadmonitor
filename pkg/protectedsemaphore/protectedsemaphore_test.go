package protectedsemaphore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		initialSize   int
		expectedSize  int
		expectedEmpty bool
	}{
		{"Positive size", 5, 5, true},
		{"Zero size", 0, 0, true},
		{"Negative size", -3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := New(tt.initialSize)
			assert.Equal(t, poolSize(tt.expectedSize), ps.poolSize)
			assert.Equal(t, tt.expectedEmpty, len(ps.semaphoreChan) == 0)
		})
	}
}

func TestIncrease(t *testing.T) {
	tests := []struct {
		name        string
		initialSize int
		increases   int
		expectFill  int
	}{
		{"Increase within bounds", 3, 2, 2},
		{"Increase to full capacity", 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := New(tt.initialSize)
			for i := 0; i < tt.increases; i++ {
				ps.Increase()
			}
			assert.Equal(t, tt.expectFill, len(ps.semaphoreChan))
		})
	}
}

func TestDecrease(t *testing.T) {
	tests := []struct {
		name        string
		initialSize int
		increases   int
		decreases   int
		expectFill  int
	}{
		{"Decrease within bounds", 3, 3, 2, 1},
		{"Decrease to empty", 3, 3, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := New(tt.initialSize)
			for i := 0; i < tt.increases; i++ {
				ps.Increase()
			}
			for i := 0; i < tt.decreases; i++ {
				ps.Decrease()
			}
			assert.Equal(t, tt.expectFill, len(ps.semaphoreChan))
		})
	}
}

func TestSetSize(t *testing.T) {
	tests := []struct {
		name        string
		initialSize int
		newSize     int
		expectSize  int
		expectFill  int
	}{
		{"Set size smaller", 5, 3, 3, 3},
		{"Set size larger", 3, 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := New(tt.initialSize)

			ps.SetSize(tt.newSize)
			assert.Equal(t, poolSize(tt.expectSize), ps.poolSize)
			assert.Equal(t, tt.expectFill, cap(ps.semaphoreChan))
		})
	}
}

func TestSetSizeDrain(t *testing.T) {
	tests := []struct {
		name        string
		initialSize int
		newSize     int
	}{
		{"Set size smaller", 5, 3},
		{"Set size larger", 3, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := New(tt.initialSize)
			ps.SetSize(tt.newSize)
			assert.Equal(t, 0, len(ps.semaphoreChan))
		})
	}
}
