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
