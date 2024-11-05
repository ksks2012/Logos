package systems

import (
	"sync"
	"time"

	"github.com/logos/ecs/entities/unit"
)

type ExperienceSystem struct {
	Unit []*unit.Unit
}

func (es *ExperienceSystem) Update() {
	// Split units into smaller chunks and process them in parallel
	// TODO: Using parameter in config file
	numWorkers := 4
	unitChunkSize := len(es.Unit) / numWorkers
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * unitChunkSize
		end := start + unitChunkSize
		if i == numWorkers-1 {
			end = len(es.Unit)
		}

		wg.Add(1)
		go func(units []*unit.Unit) {
			defer wg.Done()
			for _, e := range units {
				e.Experience.Level++
				if e.Experience.Level > 10 {
					e.Experience.Level = 1
					e.Experience.Stage += 1
				}
			}
		}(es.Unit[start:end])
	}

	// Wait for all goroutines to finish
	wg.Wait()
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
