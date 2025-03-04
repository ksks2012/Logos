package ecs_systems

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
	"gopkg.in/yaml.v2"
)

func saveToFile(data interface{}, filename string, encoderFunc func(file *os.File) error) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	filename = global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	err := os.MkdirAll(global.SaveLoadSetting.SavePath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return encoderFunc(file)
}

func SaveUnitToJSON(unit unit.Unit, filename string) error {
	return saveToFile(unit, filename, func(file *os.File) error {
		encoder := json.NewEncoder(file)
		return encoder.Encode(unit)
	})
}

func SaveUnitsToJSON(units *[]unit.Unit, filename string) error {
	return saveToFile(units, filename, func(file *os.File) error {
		encoder := json.NewEncoder(file)
		return encoder.Encode(units)
	})
}

func SaveUnitToYAML(unit unit.Unit, filename string) error {
	return saveToFile(unit, filename, func(file *os.File) error {
		encoder := yaml.NewEncoder(file)
		return encoder.Encode(unit)
	})
}

func SaveUnitsToYAML(units *[]unit.Unit, filename string) error {
	return saveToFile(units, filename, func(file *os.File) error {
		encoder := yaml.NewEncoder(file)
		return encoder.Encode(units)
	})
}

func SaveUnitState(unit unit.Unit, filename string) error {
	switch global.SaveLoadSetting.SaveFileExt {
	case ".json":
		return SaveUnitToJSON(unit, filename)
	case ".yaml", ".yml":
		return SaveUnitToYAML(unit, filename)
	default:
		return errors.New("unsupported file extension")
	}
}

func SaveUnitsState(units *[]unit.Unit, filename string) error {
	switch global.SaveLoadSetting.SaveFileExt {
	case ".json":
		return SaveUnitsToJSON(units, filename)
	case ".yaml", ".yml":
		return SaveUnitsToYAML(units, filename)
	default:
		return errors.New("unsupported file extension")
	}
}
