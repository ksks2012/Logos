package info

import (
	"fmt"

	"github.com/logos/ecs/entities/unit"
)

func ListUnits(units []unit.Unit) {
	// List all units
	fmt.Printf("| %-10s | %-10s | %-10s |\n", "Name", "Stage", "Location")
	fmt.Println("|------------|------------|------------|")
	for _, u := range units {
		fmt.Printf("| %-10s | %-10d | %-10d |\n", u.Name, u.Experience.Stage, u.LocationID)
	}
}
