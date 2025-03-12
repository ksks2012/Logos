package ecs_systems

import (
	"github.com/logos/ecs/entities"
	"github.com/logos/ecs/entities/unit"
)

func UpdateBiology(world *entities.World, biologies []*unit.Unit) {
	// TODO: Check life of unit

	// Update experience of units
	ecs_systems := ExperienceSystem{Unit: biologies}
	ecs_systems.Update()
}
