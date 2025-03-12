package system_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/ecs_systems"
	"github.com/logos/ecs/entities/unit"
)

func TestCheckLife(t *testing.T) {
	tests := []struct {
		name     string
		units    []*unit.Unit
		expected bool
	}{
		{
			name: "All units alive",
			units: []*unit.Unit{
				{Character: attributes.CharacterAttributes{Age: 20}, Experience: attributes.Experience{Stage: 1}, IsAlive: true},
				{Character: attributes.CharacterAttributes{Age: 30}, Experience: attributes.Experience{Stage: 2}, IsAlive: true},
			},
			expected: true,
		},
		{
			name: "Some units may die",
			units: []*unit.Unit{
				{Character: attributes.CharacterAttributes{Age: 100}, Experience: attributes.Experience{Stage: 0}, IsAlive: true},
				{Character: attributes.CharacterAttributes{Age: 200}, Experience: attributes.Experience{Stage: 1}, IsAlive: true},
				{Character: attributes.CharacterAttributes{Age: 400}, Experience: attributes.Experience{Stage: 2}, IsAlive: true},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			result := ecs_systems.CheckLife(tt.units)
			if result != tt.expected {
				t.Errorf("CheckLife() = %v, expected %v", result, tt.expected)
			}
			for _, u := range tt.units {
				t.Logf("Unit %v is alive: %v", u.Character.Age, u.IsAlive)
				t.Logf("Max Age: %v", 100*math.Pow(2.0, float64(u.Experience.Stage)))
				if u.Character.Age >= int(100*math.Pow(2.0, float64(u.Experience.Stage))) && u.IsAlive {
					t.Errorf("Unit with age %v should not be alive", u.Character.Age)
				}
			}
		})
	}
}
