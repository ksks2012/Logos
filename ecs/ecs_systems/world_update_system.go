package ecs_systems

import (
	"github.com/logos/ecs/entities"
	"github.com/logos/ecs/entities/unit"
)

func UpdateWorld(world *entities.World, units []*unit.Unit) {

	UpdateBiology(world, units)

}
