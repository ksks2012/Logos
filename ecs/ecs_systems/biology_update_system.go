package ecs_systems

import (
	"math"
	"math/rand"
	"time"

	"github.com/logos/ecs/entities"
	"github.com/logos/ecs/entities/unit"
)

func UpdateBiology(world *entities.World, biologies []*unit.Unit) {
	// TODO: Check life of unit
	UpdateAge(biologies)
	CheckLife(biologies)

	// Update experience of units
	experienceSystem := ExperienceSystem{Unit: biologies}
	experienceSystem.Update()
}

func calculateDeathProbability(age, maxAge, k float64) float64 {
	return 1 - math.Exp(-math.Pow(age/maxAge, k))
}

func CheckLife(units []*unit.Unit) bool {
	rand.Seed(time.Now().UnixNano())
	// NOTE: Adjust the growth rate coefficient as needed
	k := 3 + rand.Float64()*2

	for _, u := range units {
		if !u.IsAlive {
			continue
		}
		age := float64(u.Character.Age)
		// TODO: Storage in Unit structure
		maxAge := 100 * math.Pow(2, float64(u.Experience.Stage))
		deathProbability := calculateDeathProbability(age, maxAge, k)
		if rand.Float64() < deathProbability {
			u.IsAlive = false
		}
	}

	return true
}

func UpdateAge(units []*unit.Unit) {
	for _, u := range units {
		if !u.IsAlive {
			continue
		}
		u.Character.Age++
	}
}
