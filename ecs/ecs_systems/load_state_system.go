package ecs_systems

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
	"gopkg.in/yaml.v2"
)

type Decoder interface {
	Decode(v interface{}) error
}

type jsonDecoder struct {
	decoder *json.Decoder
}

func (jd *jsonDecoder) Decode(v interface{}) error {
	return jd.decoder.Decode(v)
}

type yamlDecoder struct {
	decoder *yaml.Decoder
}

func (yd *yamlDecoder) Decode(v interface{}) error {
	return yd.decoder.Decode(v)
}

func getDecoder(file *os.File, ext string) (Decoder, error) {
	switch ext {
	case ".json":
		return &jsonDecoder{decoder: json.NewDecoder(file)}, nil
	case ".yaml", ".yml":
		return &yamlDecoder{decoder: yaml.NewDecoder(file)}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}

func loadFromFile(data interface{}, filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	filename = global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder, err := getDecoder(file, global.SaveLoadSetting.SaveFileExt)
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}

func LoadUnitState(filename string) (unit.Unit, error) {
	var u unit.Unit
	err := loadFromFile(&u, filename)
	return u, err
}

func LoadUnitsState(filename string) ([]unit.Unit, error) {
	var units []unit.Unit
	err := loadFromFile(&units, filename)
	return units, err
}
