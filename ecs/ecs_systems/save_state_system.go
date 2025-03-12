package ecs_systems

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
	"gopkg.in/yaml.v2"
)

type Encoder interface {
	Encode(v interface{}) error
}

type jsonEncoder struct {
	encoder *json.Encoder
}

func (je *jsonEncoder) Encode(v interface{}) error {
	return je.encoder.Encode(v)
}

type yamlEncoder struct {
	encoder *yaml.Encoder
}

func (ye *yamlEncoder) Encode(v interface{}) error {
	return ye.encoder.Encode(v)
}

func getEncoder(file *os.File, ext string) (Encoder, error) {
	switch ext {
	case ".json":
		return &jsonEncoder{encoder: json.NewEncoder(file)}, nil
	case ".yaml", ".yml":
		return &yamlEncoder{encoder: yaml.NewEncoder(file)}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}

func saveToFile(data interface{}, filename string) error {
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

	encoder, err := getEncoder(file, global.SaveLoadSetting.SaveFileExt)
	if err != nil {
		return err
	}

	return encoder.Encode(data)
}

func SaveUnitState(unit unit.Unit, filename string) error {
	return saveToFile(unit, filename)
}

func SaveUnitsState(units []*unit.Unit, filename string) error {
	return saveToFile(units, filename)
}
