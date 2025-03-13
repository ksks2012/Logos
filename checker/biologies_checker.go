package checker

import (
	"fmt"

	"github.com/logos/ecs/entities/unit"
)

func CheckAvgDeadAge(units []*unit.Unit) {
	var sum int
	var count int
	var maxi int
	for _, u := range units {
		if !u.IsAlive {
			sum += u.Character.Age
			count++
			maxi = max(maxi, u.Character.Age)
		}
	}
	avg := sum / count
	fmt.Printf("Average age of dead units: %v\n", avg)
	fmt.Printf("Max age of dead units: %v\n", maxi)
}

func CheckReapetedID(units []*unit.Unit) {
	idMap := make(map[int64]bool)
	for _, u := range units {
		if idMap[u.ID] {
			fmt.Printf("Repeated ID: %v\n", u.ID)
		}
		idMap[u.ID] = true
	}
	fmt.Printf("ID is unique\n")
}
