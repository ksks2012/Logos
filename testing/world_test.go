package testing

import (
	"testing"

	"github.com/logos/ecs/entities"
)

func TestNewWorld(t *testing.T) {
	tests := []struct {
		name string
		x, y int
	}{
		{"10x10", 10, 10},
		{"0x0", 0, 0},
		{"1x1", 1, 1},
		{"100x100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := entities.NewWorld(tt.x, tt.y)

			if len(w.Blocks) != tt.x*tt.y {
				t.Errorf("expected %d blocks, got %d", tt.x*tt.y, len(w.Blocks))
			}
		})
	}
}
