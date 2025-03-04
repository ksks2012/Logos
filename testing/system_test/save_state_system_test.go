package system_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/ecs_systems"
	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
	"github.com/logos/pkg/setting"
	"gopkg.in/yaml.v2"
)

func TestSaveUnitState(t *testing.T) {
	tests := []struct {
		name     string
		unit     unit.Unit
		filename string
		ext      string
		wantErr  bool
	}{
		{"ValidUnitJSON", unit.Unit{
			Name:       "TestUnit",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "test_unit", ".json", false},
		{"ValidUnitYAML", unit.Unit{
			Name:       "TestUnit",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "test_unit", ".yaml", false},
		{"EmptyFilenameJSON", unit.Unit{
			Name:       "TestUnit2",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "", ".json", true},
		{"EmptyFilenameYAML", unit.Unit{
			Name:       "TestUnit2",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "", ".yaml", true},
		{"UnsupportedExtension", unit.Unit{
			Name:       "TestUnit3",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "test_unit", ".txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			global.SaveLoadSetting = &setting.SaveLoadSettingS{}
			global.SaveLoadSetting.SavePath = "./var"
			global.SaveLoadSetting.SaveFileExt = tt.ext

			err := ecs_systems.SaveUnitState(tt.unit, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUnitState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				filePath := filepath.Join(global.SaveLoadSetting.SavePath, tt.filename+"sv"+tt.ext)
				defer os.Remove(filePath)

				file, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				defer file.Close()

				var savedUnit unit.Unit
				if tt.ext == ".json" {
					if err := json.NewDecoder(file).Decode(&savedUnit); err != nil {
						t.Fatalf("failed to decode JSON: %v", err)
					}
				} else if tt.ext == ".yaml" {
					if err := yaml.NewDecoder(file).Decode(&savedUnit); err != nil {
						t.Fatalf("failed to decode YAML: %v", err)
					}
				}

				if savedUnit != tt.unit {
					t.Errorf("expected unit %v, got %v", tt.unit, savedUnit)
				}
			}
		})
	}
}

func TestSaveUnitsState(t *testing.T) {
	tests := []struct {
		name     string
		units    []unit.Unit
		filename string
		ext      string
		wantErr  bool
	}{
		{"ValidUnitsJSON", []unit.Unit{
			{
				Name:       "TestUnit1",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 1,
			},
			{
				Name:       "TestUnit2",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 2,
			},
		}, "test_units", ".json", false},
		{"ValidUnitsYAML", []unit.Unit{
			{
				Name:       "TestUnit1",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 1,
			},
			{
				Name:       "TestUnit2",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 2,
			},
		}, "test_units", ".yaml", false},
		{"EmptyFilenameJSON", []unit.Unit{
			{
				Name:       "TestUnit3",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 3,
			},
		}, "", ".json", true},
		{"EmptyFilenameYAML", []unit.Unit{
			{
				Name:       "TestUnit3",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 3,
			},
		}, "", ".yaml", true},
		{"UnsupportedExtension", []unit.Unit{
			{
				Name:       "TestUnit4",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 4,
			},
		}, "test_units", ".txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			global.SaveLoadSetting = &setting.SaveLoadSettingS{}
			global.SaveLoadSetting.SavePath = "./var"
			global.SaveLoadSetting.SaveFileExt = tt.ext

			err := ecs_systems.SaveUnitsState(&tt.units, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUnitsState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				filePath := filepath.Join(global.SaveLoadSetting.SavePath, tt.filename+"sv"+tt.ext)
				defer os.Remove(filePath)

				file, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				defer file.Close()

				var savedUnits []unit.Unit
				if tt.ext == ".json" {
					if err := json.NewDecoder(file).Decode(&savedUnits); err != nil {
						t.Fatalf("failed to decode JSON: %v", err)
					}
				} else if tt.ext == ".yaml" {
					if err := yaml.NewDecoder(file).Decode(&savedUnits); err != nil {
						t.Fatalf("failed to decode YAML: %v", err)
					}
				}

				if len(savedUnits) != len(tt.units) {
					t.Errorf("expected %d units, got %d", len(tt.units), len(savedUnits))
				}

				for i, savedUnit := range savedUnits {
					if savedUnit != tt.units[i] {
						t.Errorf("expected unit %v, got %v", tt.units[i], savedUnit)
					}
				}
			}
		})
	}
}
