package ecs_systems

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
)

func SaveUnitToJSON(unit unit.Unit, filename string) error {

	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	filename = global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	print(filename)
	err := os.MkdirAll(global.SaveLoadSetting.SavePath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(unit)
}

func SaveUnitsToJSON(units *[]unit.Unit, filename string) error {

	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	filename = global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	print(filename)
	err := os.MkdirAll(global.SaveLoadSetting.SavePath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(units)
}
