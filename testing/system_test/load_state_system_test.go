package system_test

import (
	"os"
	"path/filepath"
	"testing"

	"encoding/json"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/ecs_systems"
	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
	"github.com/logos/pkg/setting"
	"gopkg.in/yaml.v2"
)

func TestLoadFromFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		ext      string
		content  interface{}
		wantErr  bool
	}{
		{"ValidJSON", "test_unit", ".json", unit.Unit{
			Name:       "TestUnit",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, false},
		{"ValidYAML", "test_unit", ".yaml", unit.Unit{
			Name:       "TestUnit",
			Character:  attributes.CharacterAttributes{},
			Practice:   attributes.PracticeAttributes{},
			Combat:     attributes.CombatAttributes{},
			Experience: attributes.Experience{},
			LocationID: 1,
		}, false},
		{"EmptyFilename", "", ".json", unit.Unit{}, true},
		{"UnsupportedExtension", "test_unit", ".txt", unit.Unit{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			global.SaveLoadSetting = &setting.SaveLoadSettingS{}
			global.SaveLoadSetting.SavePath = "./var"
			global.SaveLoadSetting.SaveFileExt = tt.ext

			filePath := filepath.Join(global.SaveLoadSetting.SavePath, tt.filename+"sv"+tt.ext)
			if !tt.wantErr {
				file, err := os.Create(filePath)
				if err != nil {
					t.Fatalf("failed to create file: %v", err)
				}
				defer os.Remove(filePath)
				defer file.Close()

				if tt.ext == ".json" {
					if err := json.NewEncoder(file).Encode(tt.content); err != nil {
						t.Fatalf("failed to encode JSON: %v", err)
					}
				} else if tt.ext == ".yaml" {
					if err := yaml.NewEncoder(file).Encode(tt.content); err != nil {
						t.Fatalf("failed to encode YAML: %v", err)
					}
				}
			}

			var data unit.Unit
			data, err := ecs_systems.LoadUnitState(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && data != tt.content {
				t.Errorf("expected %v, got %v", tt.content, data)
			}
		})
	}
}
