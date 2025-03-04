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
)

func TestSaveUnitToJSON(t *testing.T) {
	global.SaveLoadSetting = &setting.SaveLoadSettingS{}
	global.SaveLoadSetting.SavePath = "./var"
	global.SaveLoadSetting.SaveFileExt = ".json"

	tests := []struct {
		name     string
		unit     unit.Unit
		filename string
		wantErr  bool
	}{
		{"ValidUnit", unit.Unit{
			Name:       "TestUnit",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "test_unit", false},
		{"EmptyFilename", unit.Unit{
			Name:       "TestUnit2",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ecs_systems.SaveUnitToJSON(tt.unit, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUnitToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				filePath := filepath.Join(global.SaveLoadSetting.SavePath, tt.filename+"sv.json")
				defer os.Remove(filePath)

				file, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				defer file.Close()

				var savedUnit unit.Unit
				if err := json.NewDecoder(file).Decode(&savedUnit); err != nil {
					t.Fatalf("failed to decode JSON: %v", err)
				}

				if savedUnit != tt.unit {
					t.Errorf("expected unit %v, got %v", tt.unit, savedUnit)
				}
			}
		})
	}
}

func TestSaveUnitsToJSON(t *testing.T) {
	global.SaveLoadSetting = &setting.SaveLoadSettingS{}
	global.SaveLoadSetting.SavePath = "./var"
	global.SaveLoadSetting.SaveFileExt = ".json"

	tests := []struct {
		name     string
		units    []unit.Unit
		filename string
		wantErr  bool
	}{
		{"ValidUnits", []unit.Unit{
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
		}, "test_units", false},
		{"EmptyFilename", []unit.Unit{
			{
				Name:       "TestUnit3",
				Character:  attributes.CharacterAttributes{},
				Practice:   attributes.PracticeAttributes{},
				Combat:     attributes.CombatAttributes{},
				Experience: attributes.Experience{},
				LocationID: 3,
			},
		}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ecs_systems.SaveUnitsToJSON(&tt.units, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUnitsToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				filePath := filepath.Join(global.SaveLoadSetting.SavePath, tt.filename+"sv.json")
				defer os.Remove(filePath)

				file, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				defer file.Close()

				var savedUnits []unit.Unit
				if err := json.NewDecoder(file).Decode(&savedUnits); err != nil {
					t.Fatalf("failed to decode JSON: %v", err)
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
