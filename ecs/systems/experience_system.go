package systems

import (
	"time"

	"github.com/logos/ecs/entities/unit"
)

type ExperienceSystem struct {
	Unit []*unit.Unit
}

func (es *ExperienceSystem) Update() {
	for _, e := range es.Unit {
		e.Experience.Level++
		if e.Experience.Level > 10 {
			e.Experience.Level = 1
			e.Experience.Stage += 1
		}
	}
}

// Goroutine to regularly update positions
func (es *ExperienceSystem) Start() {
	go func() {
		for {
			es.Update()
			time.Sleep(16 * time.Millisecond) // 60 updates per second
		}
	}()
}
