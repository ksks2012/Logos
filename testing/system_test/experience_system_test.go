package system_test

import (
	"testing"
	"time"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/ecs_systems"
	"github.com/logos/ecs/entities/unit"
)

func TestExperienceSystem_Update(t *testing.T) {
	tests := []struct {
		name     string
		units    []*unit.Unit
		expected []*unit.Unit
	}{
		{
			name: "Level up within stage",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 1}},
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
				{Experience: attributes.Experience{Level: 3, Stage: 1}},
			},
		},
		{
			name: "Level up and stage up",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 10, Stage: 1}},
				{Experience: attributes.Experience{Level: 1, Stage: 2}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 2}},
				{Experience: attributes.Experience{Level: 2, Stage: 2}},
			},
		},
	}

	for _, tt := range tests {
		es := &ecs_systems.ExperienceSystem{
			Unit: tt.units,
		}
		es.Update()
		for i, e := range es.Unit {
			if e.Experience.Level != tt.expected[i].Experience.Level {
				t.Errorf("expected level %d, got %d", tt.expected[i].Experience.Level, e.Experience.Level)
			}
			if e.Experience.Stage != tt.expected[i].Experience.Stage {
				t.Errorf("expected stage %d, got %d", tt.expected[i].Experience.Stage, e.Experience.Stage)
			}
		}
	}
}

func TestExperienceSystem_Update_Concurrent(t *testing.T) {
	tests := []struct {
		name     string
		units    []*unit.Unit
		expected []*unit.Unit
	}{
		{
			name: "Concurrent level up within stage",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 1}},
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
				{Experience: attributes.Experience{Level: 3, Stage: 1}},
				{Experience: attributes.Experience{Level: 4, Stage: 1}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
				{Experience: attributes.Experience{Level: 3, Stage: 1}},
				{Experience: attributes.Experience{Level: 4, Stage: 1}},
				{Experience: attributes.Experience{Level: 5, Stage: 1}},
			},
		},
		{
			name: "Concurrent level up and stage up",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 10, Stage: 1}},
				{Experience: attributes.Experience{Level: 9, Stage: 1}},
				{Experience: attributes.Experience{Level: 8, Stage: 1}},
				{Experience: attributes.Experience{Level: 7, Stage: 1}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 2}},
				{Experience: attributes.Experience{Level: 10, Stage: 1}},
				{Experience: attributes.Experience{Level: 9, Stage: 1}},
				{Experience: attributes.Experience{Level: 8, Stage: 1}},
			},
		},
	}

	for _, tt := range tests {
		es := &ecs_systems.ExperienceSystem{
			Unit: tt.units,
		}
		es.Update()
		for i, e := range es.Unit {
			if e.Experience.Level != tt.expected[i].Experience.Level {
				t.Errorf("expected level %d, got %d", tt.expected[i].Experience.Level, e.Experience.Level)
			}
			if e.Experience.Stage != tt.expected[i].Experience.Stage {
				t.Errorf("expected stage %d, got %d", tt.expected[i].Experience.Stage, e.Experience.Stage)
			}
		}
	}
}

func TestExperienceSystem_Start(t *testing.T) {
	tests := []struct {
		name     string
		units    []*unit.Unit
		expected []*unit.Unit
		interval time.Duration
	}{
		{
			name: "Regular interval update",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 1}},
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 2, Stage: 1}},
				{Experience: attributes.Experience{Level: 3, Stage: 1}},
			},
			interval: 16 * time.Millisecond,
		},
		{
			name: "Level up and stage up with interval",
			units: []*unit.Unit{
				{Experience: attributes.Experience{Level: 10, Stage: 1}},
				{Experience: attributes.Experience{Level: 9, Stage: 1}},
			},
			expected: []*unit.Unit{
				{Experience: attributes.Experience{Level: 1, Stage: 2}},
				{Experience: attributes.Experience{Level: 10, Stage: 1}},
			},
			interval: 16 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		es := &ecs_systems.ExperienceSystem{
			Unit: tt.units,
		}
		done := make(chan struct{})
		go es.Start(tt.interval, done)

		// Wait for the interval to ensure the update has been called
		time.Sleep(tt.interval + 10*time.Millisecond)
		close(done)

		for i, e := range es.Unit {
			if e.Experience.Level != tt.expected[i].Experience.Level {
				t.Errorf("expected level %d, got %d", tt.expected[i].Experience.Level, e.Experience.Level)
			}
			if e.Experience.Stage != tt.expected[i].Experience.Stage {
				t.Errorf("expected stage %d, got %d", tt.expected[i].Experience.Stage, e.Experience.Stage)
			}
		}
	}
}
