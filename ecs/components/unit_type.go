package components

import (
	"math/rand"
	"time"
)

type UnitTypeEnum int

const (
	Human UnitTypeEnum = iota
	Monster
	MaxiUnitType
)

type UnitType struct {
	Name   string       `json:"name"`
	TypeID UnitTypeEnum `json:"type_id"`
}

func NewUnitType(name string, typeID UnitTypeEnum) UnitType {
	return UnitType{
		Name:   name,
		TypeID: typeID,
	}
}

func RandomUnitType() UnitType {
	rand.Seed(time.Now().UnixNano())
	randomTypeID := UnitTypeEnum(rand.Intn(int(MaxiUnitType - 1)))
	return UnitType{
		Name:   "Random Unit Type",
		TypeID: randomTypeID,
	}
}
