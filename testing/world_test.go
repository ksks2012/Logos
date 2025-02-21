package testing

import (
	"os"
	"testing"

	"github.com/logos/ecs/entities"
	"github.com/logos/global"
	"github.com/logos/pkg/setting"
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
			w, err := entities.NewWorld(tt.x, tt.y)

			if err != nil {
				t.Logf("expected no error, got %v", err)
			} else if len(w.Blocks) != tt.x*tt.y {
				t.Errorf("expected %d blocks, got %d", tt.x*tt.y, len(w.Blocks))
			}
		})
	}
}
func TestRandomLocation(t *testing.T) {
	tests := []struct {
		name string
		x, y int
	}{
		{"10x10", 10, 10},
		{"1x1", 1, 1},
		{"100x100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := entities.NewWorld(tt.x, tt.y)
			block := w.RandomLocation()

			found := false
			for _, b := range w.Blocks {
				if b == block {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("random location block not found in world blocks")
			}
		})
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	os.Exit(code)
}

func setup() {
	// Initialize global settings
	global.WorldAttributeSetting = &setting.WorldSettingS{
		BlockSizeRange: setting.AttributeRange{
			Min: 5,
			Max: 10,
		},
		AuraRange: setting.AttributeRange{
			Min: 50,
			Max: 100,
		},
	}
}
